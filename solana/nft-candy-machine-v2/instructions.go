// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"bytes"
	"fmt"

	"github.com/sonr-io/coins/solana/base"
)

var (
	Instruction_MintNft = base.TypeID([8]byte{211, 57, 6, 167, 15, 219, 35, 251})

	Instruction_UpdateCandyMachine = base.TypeID([8]byte{243, 251, 124, 156, 211, 211, 118, 239})

	Instruction_AddConfigLines = base.TypeID([8]byte{223, 50, 224, 227, 151, 8, 115, 106})

	Instruction_InitializeCandyMachine = base.TypeID([8]byte{142, 137, 167, 107, 47, 39, 240, 124})

	Instruction_UpdateAuthority = base.TypeID([8]byte{32, 46, 64, 28, 149, 75, 243, 88})

	Instruction_WithdrawFunds = base.TypeID([8]byte{241, 36, 29, 111, 208, 31, 104, 217})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id base.TypeID) string {
	switch id {
	case Instruction_MintNft:
		return "MintNft"
	case Instruction_UpdateCandyMachine:
		return "UpdateCandyMachine"
	case Instruction_AddConfigLines:
		return "AddConfigLines"
	case Instruction_InitializeCandyMachine:
		return "InitializeCandyMachine"
	case Instruction_UpdateAuthority:
		return "UpdateAuthority"
	case Instruction_WithdrawFunds:
		return "WithdrawFunds"
	default:
		return ""
	}
}

type Instruction struct {
	base.BaseVariant
}

func (inst *Instruction) ProgramID() base.PublicKey {
	return base.MetaplexCandyMachineV2ProgramID
}

func (inst *Instruction) Accounts() (out []*base.AccountMeta) {
	return inst.Impl.(base.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := base.NewBinEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) MarshalWithEncoder(encoder *base.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}
