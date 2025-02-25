// Code generated by: `make actors-gen`. DO NOT EDIT.
package verifreg

import (
	"fmt"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-state-types/cbor"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
	builtin8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"

	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/lily/chain/actors"
	"github.com/filecoin-project/lily/chain/actors/adt"
	lotusactors "github.com/filecoin-project/lotus/chain/actors"
)

var (
	Address = builtin8.VerifiedRegistryActorAddr
	Methods = builtin8.MethodsVerifiedRegistry
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := lotusactors.GetActorMetaByCode(act.Code); ok {
		if name != actors.VerifregKey {
			return nil, fmt.Errorf("actor code is not verifreg: %s", name)
		}

		switch actors.Version(av) {

		case actors.Version8:
			return load8(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.VerifiedRegistryActorCodeID:
		return load0(store, act.Head)

	case builtin2.VerifiedRegistryActorCodeID:
		return load2(store, act.Head)

	case builtin3.VerifiedRegistryActorCodeID:
		return load3(store, act.Head)

	case builtin4.VerifiedRegistryActorCodeID:
		return load4(store, act.Head)

	case builtin5.VerifiedRegistryActorCodeID:
		return load5(store, act.Head)

	case builtin6.VerifiedRegistryActorCodeID:
		return load6(store, act.Head)

	case builtin7.VerifiedRegistryActorCodeID:
		return load7(store, act.Head)

	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Marshaler

	Code() cid.Cid
	ActorKey() string
	ActorVersion() actors.Version

	RootKey() (address.Address, error)
	VerifiedClientDataCap(address.Address) (bool, abi.StoragePower, error)
	VerifierDataCap(address.Address) (bool, abi.StoragePower, error)
	ForEachVerifier(func(addr address.Address, dcap abi.StoragePower) error) error
	ForEachClient(func(addr address.Address, dcap abi.StoragePower) error) error

	VerifiersMap() (adt.Map, error)
	VerifiersMapBitWidth() int
	VerifiersMapHashFunction() func(input []byte) []byte

	VerifiedClientsMap() (adt.Map, error)
	VerifiedClientsMapBitWidth() int
	VerifiedClientsMapHashFunction() func(input []byte) []byte
}

type VerifierInfo struct {
	Address address.Address
	DataCap abi.StoragePower
}

type VerifierChange struct {
	Before VerifierInfo
	After  VerifierInfo
}

type VerifierChanges struct {
	Added    []VerifierInfo
	Modified []VerifierChange
	Removed  []VerifierInfo
}

func AllCodes() []cid.Cid {
	return []cid.Cid{
		(&state0{}).Code(),
		(&state2{}).Code(),
		(&state3{}).Code(),
		(&state4{}).Code(),
		(&state5{}).Code(),
		(&state6{}).Code(),
		(&state7{}).Code(),
		(&state8{}).Code(),
	}
}

func VersionCodes() map[actors.Version]cid.Cid {
	return map[actors.Version]cid.Cid{
		actors.Version0: (&state0{}).Code(),
		actors.Version2: (&state2{}).Code(),
		actors.Version3: (&state3{}).Code(),
		actors.Version4: (&state4{}).Code(),
		actors.Version5: (&state5{}).Code(),
		actors.Version6: (&state6{}).Code(),
		actors.Version7: (&state7{}).Code(),
		actors.Version8: (&state8{}).Code(),
	}
}
