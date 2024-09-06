// Code generated by ent, DO NOT EDIT.

package generationmodel

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldID, id))
}

// NameInWorker applies equality check predicate on the "name_in_worker" field. It's identical to NameInWorkerEQ.
func NameInWorker(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldNameInWorker, v))
}

// ShortName applies equality check predicate on the "short_name" field. It's identical to ShortNameEQ.
func ShortName(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldShortName, v))
}

// IsActive applies equality check predicate on the "is_active" field. It's identical to IsActiveEQ.
func IsActive(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsActive, v))
}

// IsDefault applies equality check predicate on the "is_default" field. It's identical to IsDefaultEQ.
func IsDefault(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsDefault, v))
}

// IsHidden applies equality check predicate on the "is_hidden" field. It's identical to IsHiddenEQ.
func IsHidden(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsHidden, v))
}

// RunpodEndpoint applies equality check predicate on the "runpod_endpoint" field. It's identical to RunpodEndpointEQ.
func RunpodEndpoint(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldRunpodEndpoint, v))
}

// RunpodActive applies equality check predicate on the "runpod_active" field. It's identical to RunpodActiveEQ.
func RunpodActive(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldRunpodActive, v))
}

// DisplayWeight applies equality check predicate on the "display_weight" field. It's identical to DisplayWeightEQ.
func DisplayWeight(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDisplayWeight, v))
}

// DefaultSchedulerID applies equality check predicate on the "default_scheduler_id" field. It's identical to DefaultSchedulerIDEQ.
func DefaultSchedulerID(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultSchedulerID, v))
}

// DefaultWidth applies equality check predicate on the "default_width" field. It's identical to DefaultWidthEQ.
func DefaultWidth(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultWidth, v))
}

// DefaultHeight applies equality check predicate on the "default_height" field. It's identical to DefaultHeightEQ.
func DefaultHeight(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultHeight, v))
}

// DefaultInferenceSteps applies equality check predicate on the "default_inference_steps" field. It's identical to DefaultInferenceStepsEQ.
func DefaultInferenceSteps(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultInferenceSteps, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameInWorkerEQ applies the EQ predicate on the "name_in_worker" field.
func NameInWorkerEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldNameInWorker, v))
}

// NameInWorkerNEQ applies the NEQ predicate on the "name_in_worker" field.
func NameInWorkerNEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldNameInWorker, v))
}

// NameInWorkerIn applies the In predicate on the "name_in_worker" field.
func NameInWorkerIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldNameInWorker, vs...))
}

// NameInWorkerNotIn applies the NotIn predicate on the "name_in_worker" field.
func NameInWorkerNotIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldNameInWorker, vs...))
}

// NameInWorkerGT applies the GT predicate on the "name_in_worker" field.
func NameInWorkerGT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldNameInWorker, v))
}

// NameInWorkerGTE applies the GTE predicate on the "name_in_worker" field.
func NameInWorkerGTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldNameInWorker, v))
}

// NameInWorkerLT applies the LT predicate on the "name_in_worker" field.
func NameInWorkerLT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldNameInWorker, v))
}

// NameInWorkerLTE applies the LTE predicate on the "name_in_worker" field.
func NameInWorkerLTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldNameInWorker, v))
}

// NameInWorkerContains applies the Contains predicate on the "name_in_worker" field.
func NameInWorkerContains(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContains(FieldNameInWorker, v))
}

// NameInWorkerHasPrefix applies the HasPrefix predicate on the "name_in_worker" field.
func NameInWorkerHasPrefix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasPrefix(FieldNameInWorker, v))
}

// NameInWorkerHasSuffix applies the HasSuffix predicate on the "name_in_worker" field.
func NameInWorkerHasSuffix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasSuffix(FieldNameInWorker, v))
}

// NameInWorkerEqualFold applies the EqualFold predicate on the "name_in_worker" field.
func NameInWorkerEqualFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEqualFold(FieldNameInWorker, v))
}

// NameInWorkerContainsFold applies the ContainsFold predicate on the "name_in_worker" field.
func NameInWorkerContainsFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContainsFold(FieldNameInWorker, v))
}

// ShortNameEQ applies the EQ predicate on the "short_name" field.
func ShortNameEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldShortName, v))
}

// ShortNameNEQ applies the NEQ predicate on the "short_name" field.
func ShortNameNEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldShortName, v))
}

// ShortNameIn applies the In predicate on the "short_name" field.
func ShortNameIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldShortName, vs...))
}

// ShortNameNotIn applies the NotIn predicate on the "short_name" field.
func ShortNameNotIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldShortName, vs...))
}

// ShortNameGT applies the GT predicate on the "short_name" field.
func ShortNameGT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldShortName, v))
}

// ShortNameGTE applies the GTE predicate on the "short_name" field.
func ShortNameGTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldShortName, v))
}

// ShortNameLT applies the LT predicate on the "short_name" field.
func ShortNameLT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldShortName, v))
}

// ShortNameLTE applies the LTE predicate on the "short_name" field.
func ShortNameLTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldShortName, v))
}

// ShortNameContains applies the Contains predicate on the "short_name" field.
func ShortNameContains(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContains(FieldShortName, v))
}

// ShortNameHasPrefix applies the HasPrefix predicate on the "short_name" field.
func ShortNameHasPrefix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasPrefix(FieldShortName, v))
}

// ShortNameHasSuffix applies the HasSuffix predicate on the "short_name" field.
func ShortNameHasSuffix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasSuffix(FieldShortName, v))
}

// ShortNameEqualFold applies the EqualFold predicate on the "short_name" field.
func ShortNameEqualFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEqualFold(FieldShortName, v))
}

// ShortNameContainsFold applies the ContainsFold predicate on the "short_name" field.
func ShortNameContainsFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContainsFold(FieldShortName, v))
}

// IsActiveEQ applies the EQ predicate on the "is_active" field.
func IsActiveEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsActive, v))
}

// IsActiveNEQ applies the NEQ predicate on the "is_active" field.
func IsActiveNEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldIsActive, v))
}

// IsDefaultEQ applies the EQ predicate on the "is_default" field.
func IsDefaultEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsDefault, v))
}

// IsDefaultNEQ applies the NEQ predicate on the "is_default" field.
func IsDefaultNEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldIsDefault, v))
}

// IsHiddenEQ applies the EQ predicate on the "is_hidden" field.
func IsHiddenEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldIsHidden, v))
}

// IsHiddenNEQ applies the NEQ predicate on the "is_hidden" field.
func IsHiddenNEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldIsHidden, v))
}

// RunpodEndpointEQ applies the EQ predicate on the "runpod_endpoint" field.
func RunpodEndpointEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldRunpodEndpoint, v))
}

// RunpodEndpointNEQ applies the NEQ predicate on the "runpod_endpoint" field.
func RunpodEndpointNEQ(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldRunpodEndpoint, v))
}

// RunpodEndpointIn applies the In predicate on the "runpod_endpoint" field.
func RunpodEndpointIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldRunpodEndpoint, vs...))
}

// RunpodEndpointNotIn applies the NotIn predicate on the "runpod_endpoint" field.
func RunpodEndpointNotIn(vs ...string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldRunpodEndpoint, vs...))
}

// RunpodEndpointGT applies the GT predicate on the "runpod_endpoint" field.
func RunpodEndpointGT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldRunpodEndpoint, v))
}

// RunpodEndpointGTE applies the GTE predicate on the "runpod_endpoint" field.
func RunpodEndpointGTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldRunpodEndpoint, v))
}

// RunpodEndpointLT applies the LT predicate on the "runpod_endpoint" field.
func RunpodEndpointLT(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldRunpodEndpoint, v))
}

// RunpodEndpointLTE applies the LTE predicate on the "runpod_endpoint" field.
func RunpodEndpointLTE(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldRunpodEndpoint, v))
}

// RunpodEndpointContains applies the Contains predicate on the "runpod_endpoint" field.
func RunpodEndpointContains(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContains(FieldRunpodEndpoint, v))
}

// RunpodEndpointHasPrefix applies the HasPrefix predicate on the "runpod_endpoint" field.
func RunpodEndpointHasPrefix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasPrefix(FieldRunpodEndpoint, v))
}

// RunpodEndpointHasSuffix applies the HasSuffix predicate on the "runpod_endpoint" field.
func RunpodEndpointHasSuffix(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldHasSuffix(FieldRunpodEndpoint, v))
}

// RunpodEndpointIsNil applies the IsNil predicate on the "runpod_endpoint" field.
func RunpodEndpointIsNil() predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIsNull(FieldRunpodEndpoint))
}

// RunpodEndpointNotNil applies the NotNil predicate on the "runpod_endpoint" field.
func RunpodEndpointNotNil() predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotNull(FieldRunpodEndpoint))
}

// RunpodEndpointEqualFold applies the EqualFold predicate on the "runpod_endpoint" field.
func RunpodEndpointEqualFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEqualFold(FieldRunpodEndpoint, v))
}

// RunpodEndpointContainsFold applies the ContainsFold predicate on the "runpod_endpoint" field.
func RunpodEndpointContainsFold(v string) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldContainsFold(FieldRunpodEndpoint, v))
}

// RunpodActiveEQ applies the EQ predicate on the "runpod_active" field.
func RunpodActiveEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldRunpodActive, v))
}

// RunpodActiveNEQ applies the NEQ predicate on the "runpod_active" field.
func RunpodActiveNEQ(v bool) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldRunpodActive, v))
}

// DisplayWeightEQ applies the EQ predicate on the "display_weight" field.
func DisplayWeightEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDisplayWeight, v))
}

// DisplayWeightNEQ applies the NEQ predicate on the "display_weight" field.
func DisplayWeightNEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldDisplayWeight, v))
}

// DisplayWeightIn applies the In predicate on the "display_weight" field.
func DisplayWeightIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldDisplayWeight, vs...))
}

// DisplayWeightNotIn applies the NotIn predicate on the "display_weight" field.
func DisplayWeightNotIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldDisplayWeight, vs...))
}

// DisplayWeightGT applies the GT predicate on the "display_weight" field.
func DisplayWeightGT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldDisplayWeight, v))
}

// DisplayWeightGTE applies the GTE predicate on the "display_weight" field.
func DisplayWeightGTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldDisplayWeight, v))
}

// DisplayWeightLT applies the LT predicate on the "display_weight" field.
func DisplayWeightLT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldDisplayWeight, v))
}

// DisplayWeightLTE applies the LTE predicate on the "display_weight" field.
func DisplayWeightLTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldDisplayWeight, v))
}

// DefaultSchedulerIDEQ applies the EQ predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDEQ(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDNEQ applies the NEQ predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDNEQ(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDIn applies the In predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDIn(vs ...uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldDefaultSchedulerID, vs...))
}

// DefaultSchedulerIDNotIn applies the NotIn predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDNotIn(vs ...uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldDefaultSchedulerID, vs...))
}

// DefaultSchedulerIDGT applies the GT predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDGT(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDGTE applies the GTE predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDGTE(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDLT applies the LT predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDLT(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDLTE applies the LTE predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDLTE(v uuid.UUID) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldDefaultSchedulerID, v))
}

// DefaultSchedulerIDIsNil applies the IsNil predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDIsNil() predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIsNull(FieldDefaultSchedulerID))
}

// DefaultSchedulerIDNotNil applies the NotNil predicate on the "default_scheduler_id" field.
func DefaultSchedulerIDNotNil() predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotNull(FieldDefaultSchedulerID))
}

// DefaultWidthEQ applies the EQ predicate on the "default_width" field.
func DefaultWidthEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultWidth, v))
}

// DefaultWidthNEQ applies the NEQ predicate on the "default_width" field.
func DefaultWidthNEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldDefaultWidth, v))
}

// DefaultWidthIn applies the In predicate on the "default_width" field.
func DefaultWidthIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldDefaultWidth, vs...))
}

// DefaultWidthNotIn applies the NotIn predicate on the "default_width" field.
func DefaultWidthNotIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldDefaultWidth, vs...))
}

// DefaultWidthGT applies the GT predicate on the "default_width" field.
func DefaultWidthGT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldDefaultWidth, v))
}

// DefaultWidthGTE applies the GTE predicate on the "default_width" field.
func DefaultWidthGTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldDefaultWidth, v))
}

// DefaultWidthLT applies the LT predicate on the "default_width" field.
func DefaultWidthLT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldDefaultWidth, v))
}

// DefaultWidthLTE applies the LTE predicate on the "default_width" field.
func DefaultWidthLTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldDefaultWidth, v))
}

// DefaultHeightEQ applies the EQ predicate on the "default_height" field.
func DefaultHeightEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultHeight, v))
}

// DefaultHeightNEQ applies the NEQ predicate on the "default_height" field.
func DefaultHeightNEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldDefaultHeight, v))
}

// DefaultHeightIn applies the In predicate on the "default_height" field.
func DefaultHeightIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldDefaultHeight, vs...))
}

// DefaultHeightNotIn applies the NotIn predicate on the "default_height" field.
func DefaultHeightNotIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldDefaultHeight, vs...))
}

// DefaultHeightGT applies the GT predicate on the "default_height" field.
func DefaultHeightGT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldDefaultHeight, v))
}

// DefaultHeightGTE applies the GTE predicate on the "default_height" field.
func DefaultHeightGTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldDefaultHeight, v))
}

// DefaultHeightLT applies the LT predicate on the "default_height" field.
func DefaultHeightLT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldDefaultHeight, v))
}

// DefaultHeightLTE applies the LTE predicate on the "default_height" field.
func DefaultHeightLTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldDefaultHeight, v))
}

// DefaultInferenceStepsEQ applies the EQ predicate on the "default_inference_steps" field.
func DefaultInferenceStepsEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldDefaultInferenceSteps, v))
}

// DefaultInferenceStepsNEQ applies the NEQ predicate on the "default_inference_steps" field.
func DefaultInferenceStepsNEQ(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldDefaultInferenceSteps, v))
}

// DefaultInferenceStepsIn applies the In predicate on the "default_inference_steps" field.
func DefaultInferenceStepsIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldDefaultInferenceSteps, vs...))
}

// DefaultInferenceStepsNotIn applies the NotIn predicate on the "default_inference_steps" field.
func DefaultInferenceStepsNotIn(vs ...int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldDefaultInferenceSteps, vs...))
}

// DefaultInferenceStepsGT applies the GT predicate on the "default_inference_steps" field.
func DefaultInferenceStepsGT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldDefaultInferenceSteps, v))
}

// DefaultInferenceStepsGTE applies the GTE predicate on the "default_inference_steps" field.
func DefaultInferenceStepsGTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldDefaultInferenceSteps, v))
}

// DefaultInferenceStepsLT applies the LT predicate on the "default_inference_steps" field.
func DefaultInferenceStepsLT(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldDefaultInferenceSteps, v))
}

// DefaultInferenceStepsLTE applies the LTE predicate on the "default_inference_steps" field.
func DefaultInferenceStepsLTE(v int32) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldDefaultInferenceSteps, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GenerationModel {
	return predicate.GenerationModel(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasGenerations applies the HasEdge predicate on the "generations" edge.
func HasGenerations() predicate.GenerationModel {
	return predicate.GenerationModel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GenerationsTable, GenerationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenerationsWith applies the HasEdge predicate on the "generations" edge with a given conditions (other predicates).
func HasGenerationsWith(preds ...predicate.Generation) predicate.GenerationModel {
	return predicate.GenerationModel(func(s *sql.Selector) {
		step := newGenerationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSchedulers applies the HasEdge predicate on the "schedulers" edge.
func HasSchedulers() predicate.GenerationModel {
	return predicate.GenerationModel(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SchedulersTable, SchedulersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSchedulersWith applies the HasEdge predicate on the "schedulers" edge with a given conditions (other predicates).
func HasSchedulersWith(preds ...predicate.Scheduler) predicate.GenerationModel {
	return predicate.GenerationModel(func(s *sql.Selector) {
		step := newSchedulersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GenerationModel) predicate.GenerationModel {
	return predicate.GenerationModel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GenerationModel) predicate.GenerationModel {
	return predicate.GenerationModel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GenerationModel) predicate.GenerationModel {
	return predicate.GenerationModel(sql.NotPredicates(p))
}
