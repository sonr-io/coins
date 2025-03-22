// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine_v2

import (
	"errors"

	"github.com/sonr-io/coins/solana/base"
)

// UpdateAuthority is the `updateAuthority` instruction.
type UpdateAuthority struct {
	NewAuthority *base.PublicKey `bin:"optional"`

	// [0] = [WRITE] candyMachine
	//
	// [1] = [SIGNER] authority
	//
	// [2] = [] wallet
	base.AccountMetaSlice `bin:"-"`
}

// NewUpdateAuthorityInstructionBuilder creates a new `UpdateAuthority` instruction builder.
func NewUpdateAuthorityInstructionBuilder() *UpdateAuthority {
	nd := &UpdateAuthority{
		AccountMetaSlice: make(base.AccountMetaSlice, 3),
	}
	return nd
}

// SetNewAuthority sets the "newAuthority" parameter.
func (inst *UpdateAuthority) SetNewAuthority(newAuthority base.PublicKey) *UpdateAuthority {
	inst.NewAuthority = &newAuthority
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *UpdateAuthority) SetCandyMachineAccount(candyMachine base.PublicKey) *UpdateAuthority {
	inst.AccountMetaSlice[0] = base.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *UpdateAuthority) GetCandyMachineAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" account.
func (inst *UpdateAuthority) SetAuthorityAccount(authority base.PublicKey) *UpdateAuthority {
	inst.AccountMetaSlice[1] = base.Meta(authority).SIGNER()
	return inst
}

// GetAuthorityAccount gets the "authority" account.
func (inst *UpdateAuthority) GetAuthorityAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetWalletAccount sets the "wallet" account.
func (inst *UpdateAuthority) SetWalletAccount(wallet base.PublicKey) *UpdateAuthority {
	inst.AccountMetaSlice[2] = base.Meta(wallet)
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *UpdateAuthority) GetWalletAccount() *base.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

func (inst UpdateAuthority) Build() *Instruction {
	return &Instruction{BaseVariant: base.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateAuthority,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateAuthority) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateAuthority) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Authority is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Wallet is not set")
		}
	}
	return nil
}

func (obj UpdateAuthority) MarshalWithEncoder(encoder *base.Encoder) (err error) {
	// Serialize `NewAuthority` param (optional):
	{
		if obj.NewAuthority == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.NewAuthority)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewUpdateAuthorityInstruction declares a new UpdateAuthority instruction with the provided parameters and accounts.
func NewUpdateAuthorityInstruction(
	// Parameters:
	newAuthority base.PublicKey,
	// Accounts:
	candyMachine base.PublicKey,
	authority base.PublicKey,
	wallet base.PublicKey) *UpdateAuthority {
	return NewUpdateAuthorityInstructionBuilder().
		SetNewAuthority(newAuthority).
		SetCandyMachineAccount(candyMachine).
		SetAuthorityAccount(authority).
		SetWalletAccount(wallet)
}
