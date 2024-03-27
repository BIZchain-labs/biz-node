// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://af6a3f6693c0cac25d512d19182b8934c35ece74d65e4e47b4881cac695dda714d4554b7b5987cc5a183de6622a1ae525cd854452f0367025b0acf8036bee5b6@38.180.106.246:30000",
    "enode://bf598ad7bb2008171d81145996cc821daaeaeb9f97d2786728d0daa6ef1d90cb3eaf1c99efdb2dff87f0960c7d5cd00fc429ff8eff06a95dd05c6555b1b2fcfc@38.180.106.246:30361",
    "enode://b43d5cb5f86e104714e3a5f67273ffe08b7ee35ccecad2727a59762a2e0ff576394ce82ee801d506bccee70d837c109fa9a59e359d427a8ba05d39d001aabf89@38.180.106.246:30312",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
