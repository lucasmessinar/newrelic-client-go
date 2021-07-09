// Code generated by tutone: DO NOT EDIT
package installevents

import "context"

// Creates a new recipe event.
func (a *Installevents) CreateRecipeEvent(
	accountID int,
	status RecipeStatus,
) (*RecipeEvent, error) {
	return a.CreateRecipeEventWithContext(context.Background(),
		accountID,
		status,
	)
}

// Creates a new recipe event.
func (a *Installevents) CreateRecipeEventWithContext(
	ctx context.Context,
	accountID int,
	status RecipeStatus,
) (*RecipeEvent, error) {

	resp := CreateRecipeEventQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"status":    status,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, CreateRecipeEventMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.RecipeEvent, nil
}

type CreateRecipeEventQueryResponse struct {
	RecipeEvent RecipeEvent `json:"CreateRecipeEvent"`
}

const CreateRecipeEventMutation = `mutation(
	$accountId: Int!,
	$status: RecipeStatus!,
) { installationCreateRecipeEvent(
	accountId: $accountId,
	status: $status,
) {
	cliVersion
	complete
	displayName
	entityGuid
	error {
		details
		message
	}
	hostName
	kernelArch
	kernelVersion
	logFilePath
	name
	os
	platform
	platformFamily
	platformVersion
	redirectUrl
	status
	targetedInstall
	timestamp
	validationDurationMilliseconds
} }`