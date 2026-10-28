// TEMPORARY playground: explore buttons and rotaries. Do not commit.
//
//	go run ./controls                              # list controls grouped by device
//	go run ./controls -json                        # raw JSON of buttons and rotaries
//	go run ./controls -set <buttonId> -interval 800  # try UpdateButton (repeat_interval, ms)
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sort"

	"github.com/openhue/openhue-go"
)

func main() {
	raw := flag.Bool("json", false, "print raw JSON")
	setID := flag.String("set", "", "button ID to update")
	interval := flag.Int("interval", 0, "repeat_interval in ms (with -set)")
	behaviors := flag.Bool("behaviors", false, "dump behavior scripts and instances (button -> scene assignments)")
	flag.Parse()

	ip, key := openhue.LoadConfNoError()
	ctx := context.Background()

	home, err := openhue.NewHome(ip, key)
	openhue.CheckErr(err)

	// Home has no rotary/update-button wrappers yet: use the generated client directly.
	// Throwaway only: TLS verification is skipped.
	api, err := openhue.NewClientWithResponses("https://"+ip,
		openhue.WithHTTPClient(&http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}}),
		openhue.WithRequestEditorFn(func(_ context.Context, r *http.Request) error {
			r.Header.Set("hue-application-key", key)
			return nil
		}))
	openhue.CheckErr(err)

	if *behaviors {
		scripts, err := api.GetBehaviorScriptsWithResponse(ctx)
		openhue.CheckErr(err)
		names := map[string]string{}
		if scripts.JSON200 != nil {
			for _, sc := range scripts.JSON200.Data {
				n := ""
				if sc.Metadata != nil && sc.Metadata.Name != nil {
					n = *sc.Metadata.Name
				}
				names[sc.Id] = n
			}
		}
		inst, err := api.GetBehaviorInstancesWithResponse(ctx)
		openhue.CheckErr(err)
		if inst.JSON200 == nil {
			fmt.Printf("instances: HTTP %d %s\n", inst.StatusCode(), inst.Body)
			return
		}
		for _, i := range inst.JSON200.Data {
			sid, name := "", ""
			if i.ScriptId != nil {
				sid = *i.ScriptId
			}
			if i.Metadata != nil && i.Metadata.Name != nil {
				name = *i.Metadata.Name
			}
			cfg, _ := json.MarshalIndent(i.Configuration, "", "  ")
			fmt.Printf("== %s [%s] script=%s (%s)\n%s\n", name, i.Id, names[sid], sid, cfg)
		}
		return
	}

	if *setID != "" {
		resp, err := api.UpdateButtonWithResponse(ctx, *setID, openhue.UpdateButtonJSONRequestBody{
			Button: &struct {
				RepeatInterval *int `json:"repeat_interval,omitempty"`
			}{RepeatInterval: interval},
		})
		openhue.CheckErr(err)
		fmt.Printf("PUT button %s -> HTTP %d\n%s\n", *setID, resp.StatusCode(), resp.Body)
		return
	}

	devices, err := home.GetDevices(ctx)
	openhue.CheckErr(err)
	buttons, err := home.GetButtons(ctx)
	openhue.CheckErr(err)
	rotResp, err := api.GetRelativeRotariesWithResponse(ctx)
	openhue.CheckErr(err)
	if rotResp.JSON200 == nil {
		fmt.Printf("rotaries: HTTP %d %s\n", rotResp.StatusCode(), rotResp.Body)
		return
	}
	rotaries := rotResp.JSON200.Data

	if *raw {
		out, _ := json.MarshalIndent(map[string]any{"buttons": buttons, "rotaries": rotaries}, "", "  ")
		fmt.Println(string(out))
		return
	}

	type line struct {
		control int
		text    string
	}
	byDevice := map[string][]line{}
	num := func(p *int) int {
		if p == nil {
			return -1
		}
		return *p
	}

	for _, b := range buttons {
		ev, repeat := "-", "-"
		if b.Button != nil {
			if b.Button.ButtonReport != nil && b.Button.ButtonReport.Event != nil {
				ev = string(*b.Button.ButtonReport.Event)
			}
			if b.Button.RepeatInterval != nil {
				repeat = fmt.Sprintf("%dms", *b.Button.RepeatInterval)
			}
		}
		c := 0
		if b.Metadata != nil {
			c = num(b.Metadata.ControlId)
		}
		byDevice[b.Owner.Rid] = append(byDevice[b.Owner.Rid], line{c,
			fmt.Sprintf("button  #%d  %s  last=%s repeat=%s", c, b.Id, ev, repeat)})
	}
	for _, r := range rotaries {
		ev := "-"
		if r.RelativeRotary != nil && r.RelativeRotary.RotaryReport != nil && r.RelativeRotary.RotaryReport.Action != nil {
			ev = string(*r.RelativeRotary.RotaryReport.Action)
		}
		c := 0
		if r.Metadata != nil {
			c = num(r.Metadata.ControlId)
		}
		byDevice[r.Owner.Rid] = append(byDevice[r.Owner.Rid], line{c,
			fmt.Sprintf("rotary  #%d  %s  last=%s", c, r.Id, ev)})
	}

	for devID, lines := range byDevice {
		name := devID
		if d, ok := devices[devID]; ok {
			name = d.Metadata.Name
		}
		fmt.Printf("%s (%s)\n", name, devID)
		sort.Slice(lines, func(i, j int) bool { return lines[i].control < lines[j].control })
		for _, l := range lines {
			fmt.Println("   ", l.text)
		}
	}
	fmt.Printf("\n%d buttons, %d rotaries\n", len(buttons), len(rotaries))
}
