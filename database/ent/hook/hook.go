// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/stablecog/sc-go/database/ent"
)

// The ApiTokenFunc type is an adapter to allow the use of ordinary
// function as ApiToken mutator.
type ApiTokenFunc func(context.Context, *ent.ApiTokenMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ApiTokenFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ApiTokenMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ApiTokenMutation", m)
}

// The AuthClientFunc type is an adapter to allow the use of ordinary
// function as AuthClient mutator.
type AuthClientFunc func(context.Context, *ent.AuthClientMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuthClientFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AuthClientMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuthClientMutation", m)
}

// The BannedWordsFunc type is an adapter to allow the use of ordinary
// function as BannedWords mutator.
type BannedWordsFunc func(context.Context, *ent.BannedWordsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BannedWordsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.BannedWordsMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BannedWordsMutation", m)
}

// The CreditFunc type is an adapter to allow the use of ordinary
// function as Credit mutator.
type CreditFunc func(context.Context, *ent.CreditMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CreditFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CreditMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CreditMutation", m)
}

// The CreditTypeFunc type is an adapter to allow the use of ordinary
// function as CreditType mutator.
type CreditTypeFunc func(context.Context, *ent.CreditTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CreditTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CreditTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CreditTypeMutation", m)
}

// The DeviceInfoFunc type is an adapter to allow the use of ordinary
// function as DeviceInfo mutator.
type DeviceInfoFunc func(context.Context, *ent.DeviceInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DeviceInfoMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceInfoMutation", m)
}

// The DisposableEmailFunc type is an adapter to allow the use of ordinary
// function as DisposableEmail mutator.
type DisposableEmailFunc func(context.Context, *ent.DisposableEmailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DisposableEmailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.DisposableEmailMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DisposableEmailMutation", m)
}

// The GenerationFunc type is an adapter to allow the use of ordinary
// function as Generation mutator.
type GenerationFunc func(context.Context, *ent.GenerationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenerationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GenerationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenerationMutation", m)
}

// The GenerationModelFunc type is an adapter to allow the use of ordinary
// function as GenerationModel mutator.
type GenerationModelFunc func(context.Context, *ent.GenerationModelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenerationModelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GenerationModelMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenerationModelMutation", m)
}

// The GenerationOutputFunc type is an adapter to allow the use of ordinary
// function as GenerationOutput mutator.
type GenerationOutputFunc func(context.Context, *ent.GenerationOutputMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenerationOutputFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GenerationOutputMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenerationOutputMutation", m)
}

// The GenerationOutputLikeFunc type is an adapter to allow the use of ordinary
// function as GenerationOutputLike mutator.
type GenerationOutputLikeFunc func(context.Context, *ent.GenerationOutputLikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenerationOutputLikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.GenerationOutputLikeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenerationOutputLikeMutation", m)
}

// The IPBlackListFunc type is an adapter to allow the use of ordinary
// function as IPBlackList mutator.
type IPBlackListFunc func(context.Context, *ent.IPBlackListMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f IPBlackListFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.IPBlackListMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.IPBlackListMutation", m)
}

// The MqLogFunc type is an adapter to allow the use of ordinary
// function as MqLog mutator.
type MqLogFunc func(context.Context, *ent.MqLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MqLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.MqLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MqLogMutation", m)
}

// The NegativePromptFunc type is an adapter to allow the use of ordinary
// function as NegativePrompt mutator.
type NegativePromptFunc func(context.Context, *ent.NegativePromptMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NegativePromptFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.NegativePromptMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NegativePromptMutation", m)
}

// The PromptFunc type is an adapter to allow the use of ordinary
// function as Prompt mutator.
type PromptFunc func(context.Context, *ent.PromptMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PromptFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.PromptMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PromptMutation", m)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.RoleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
}

// The SchedulerFunc type is an adapter to allow the use of ordinary
// function as Scheduler mutator.
type SchedulerFunc func(context.Context, *ent.SchedulerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SchedulerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SchedulerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SchedulerMutation", m)
}

// The ThumbmarkIdBlackListFunc type is an adapter to allow the use of ordinary
// function as ThumbmarkIdBlackList mutator.
type ThumbmarkIdBlackListFunc func(context.Context, *ent.ThumbmarkIdBlackListMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ThumbmarkIdBlackListFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ThumbmarkIdBlackListMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ThumbmarkIdBlackListMutation", m)
}

// The TipLogFunc type is an adapter to allow the use of ordinary
// function as TipLog mutator.
type TipLogFunc func(context.Context, *ent.TipLogMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TipLogFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.TipLogMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TipLogMutation", m)
}

// The UpscaleFunc type is an adapter to allow the use of ordinary
// function as Upscale mutator.
type UpscaleFunc func(context.Context, *ent.UpscaleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UpscaleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UpscaleMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UpscaleMutation", m)
}

// The UpscaleModelFunc type is an adapter to allow the use of ordinary
// function as UpscaleModel mutator.
type UpscaleModelFunc func(context.Context, *ent.UpscaleModelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UpscaleModelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UpscaleModelMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UpscaleModelMutation", m)
}

// The UpscaleOutputFunc type is an adapter to allow the use of ordinary
// function as UpscaleOutput mutator.
type UpscaleOutputFunc func(context.Context, *ent.UpscaleOutputMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UpscaleOutputFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UpscaleOutputMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UpscaleOutputMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UsernameBlacklistFunc type is an adapter to allow the use of ordinary
// function as UsernameBlacklist mutator.
type UsernameBlacklistFunc func(context.Context, *ent.UsernameBlacklistMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UsernameBlacklistFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UsernameBlacklistMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UsernameBlacklistMutation", m)
}

// The VoiceoverFunc type is an adapter to allow the use of ordinary
// function as Voiceover mutator.
type VoiceoverFunc func(context.Context, *ent.VoiceoverMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VoiceoverFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VoiceoverMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VoiceoverMutation", m)
}

// The VoiceoverModelFunc type is an adapter to allow the use of ordinary
// function as VoiceoverModel mutator.
type VoiceoverModelFunc func(context.Context, *ent.VoiceoverModelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VoiceoverModelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VoiceoverModelMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VoiceoverModelMutation", m)
}

// The VoiceoverOutputFunc type is an adapter to allow the use of ordinary
// function as VoiceoverOutput mutator.
type VoiceoverOutputFunc func(context.Context, *ent.VoiceoverOutputMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VoiceoverOutputFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VoiceoverOutputMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VoiceoverOutputMutation", m)
}

// The VoiceoverSpeakerFunc type is an adapter to allow the use of ordinary
// function as VoiceoverSpeaker mutator.
type VoiceoverSpeakerFunc func(context.Context, *ent.VoiceoverSpeakerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VoiceoverSpeakerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.VoiceoverSpeakerMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VoiceoverSpeakerMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
