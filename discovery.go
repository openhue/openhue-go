package openhue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grandcat/zeroconf"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	bridgeService  = "_hue._tcp"
	discoveryUrl   = "https://discovery.meethue.com"
	// maxDiscoveryBody caps the size of the discovery service response we are willing to read.
	maxDiscoveryBody = 1 << 20
)

type discOpt func(b *BridgeDiscovery)

type BridgeDiscovery struct {
	timeout                   time.Duration
	allowUrlDiscoveryFallback bool
}

func (e BridgeDiscoveryError) Error() string {
	return string(e)
}

type BridgeInfo struct {
	Instance  string
	HostName  string
	IpAddress string
}

func (b *BridgeInfo) String() string {
	return fmt.Sprintf("Bridge{instance: \"%s\", host: \"%s\", ip: \"%s\"}", b.Instance, b.HostName, b.IpAddress)
}

type BridgeDiscoveryError string

const (
	TimeoutError    BridgeDiscoveryError = "discovery via mDNS timeout"
	NotFoundError   BridgeDiscoveryError = "no bridge found"
	TooManyAttempts BridgeDiscoveryError = "too many attempts to discover the bridge via URL"
)

func NewBridgeDiscovery(opts ...discOpt) *BridgeDiscovery {
	bd := &BridgeDiscovery{
		timeout:                   defaultTimeout,
		allowUrlDiscoveryFallback: true,
	}
	for _, o := range opts {
		o(bd)
	}
	return bd
}

func (d *BridgeDiscovery) Discover() (*BridgeInfo, error) {

	bridgeInfo, err := mDNSDiscovery(d.timeout)
	if err != nil {
		if d.allowUrlDiscoveryFallback {
			bridgeInfo, err = urlDiscovery(d.timeout)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return bridgeInfo, nil
}

func mDNSDiscovery(timeout time.Duration) (*BridgeInfo, error) {
	resolver, err := zeroconf.NewResolver()
	if err != nil {
		return nil, err
	}

	entries := make(chan *zeroconf.ServiceEntry)
	foundEntry := make(chan *zeroconf.ServiceEntry, 1)
	var entry *zeroconf.ServiceEntry

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := resolver.Browse(ctx, bridgeService, "local", entries); err != nil {
		return nil, err
	}

	go func(results <-chan *zeroconf.ServiceEntry, foundEntry chan *zeroconf.ServiceEntry) {
		for e := range results {
			if len(e.AddrIPv4) > 0 {
				select {
				case foundEntry <- e:
					cancel()
				default: // an entry has already been reported, drain the remaining results
				}
			}
		}
	}(entries, foundEntry)

	select {
	// we only expect one hub to be found for now
	case entry = <-foundEntry:
	case <-ctx.Done():
		return nil, TimeoutError
	}

	if entry == nil {
		return nil, NotFoundError
	}

	return &BridgeInfo{
		Instance:  strings.ReplaceAll(entry.Instance, "\\", ""),
		HostName:  entry.HostName,
		IpAddress: entry.AddrIPv4[0].String(),
	}, nil
}

func urlDiscovery(timeout time.Duration) (*BridgeInfo, error) {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(discoveryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, TooManyAttempts
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("discovery service returned unexpected status %s", resp.Status)
	}

	var bridges []struct {
		ID                string `json:"id"`
		InternalIPAddress string `json:"internalipaddress"`
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxDiscoveryBody)).Decode(&bridges); err != nil {
		return nil, err
	}

	if len(bridges) == 0 || bridges[0].InternalIPAddress == "" {
		return nil, NotFoundError
	}

	return &BridgeInfo{
		Instance:  "N/A",
		HostName:  bridges[0].ID,
		IpAddress: bridges[0].InternalIPAddress,
	}, nil
}

// WithTimeout specifies that timeout value for the Bridge Discovery. Default is 5 seconds.
func WithTimeout(timeout time.Duration) discOpt {
	return func(b *BridgeDiscovery) {
		b.timeout = timeout
	}
}

// WithDisabledUrlDiscovery allows disabling the URL discovery process in case the mDNS one failed.
func WithDisabledUrlDiscovery() discOpt {
	return func(b *BridgeDiscovery) {
		b.allowUrlDiscoveryFallback = false
	}
}
