package goladok3

import (
	"context"
	"fmt"

	"github.com/masv3971/goladok3/ladoktypes"
)

// CheckPermission compare ladok permissions with ps
func (c *Client) CheckPermission(ctx context.Context, myPermissions Permissions) error {
	ctx, span := c.tp.Start(ctx, "goladok3.CheckPermission")
	defer span.End()

	var (
		e             = &Errors{}
		internalError = []ladoktypes.InternalError{}
	)

	egna, _, err := c.Kataloginformation.GetAnvandarbehorighetEgna(ctx)
	if err != nil {
		return err
	}

	if len(egna.Anvandarbehorighet) < 1 {
		return ErrNotSufficientPermissions
	}

	ladokProfile, _, err := c.Kataloginformation.GetBehorighetsprofil(ctx, &GetBehorighetsprofilerReq{UID: egna.Anvandarbehorighet[0].BehorighetsprofilRef.UID})
	if err != nil {
		return err
	}
	permissions, err := c.permissionUnify(ctx, *ladokProfile, myPermissions)
	if err != nil {
		return err
	}

	for permissionID, data := range permissions {
		myPermission, ok := data["my"]
		if !ok {
			// continue if ladok has permission not specified in "my"
			continue
		}

		ladokPermission, ok := data["ladok"]
		if !ok {
			// Ladok does not have the required permission
			internalError = append(internalError, ladoktypes.InternalError{
				Msg:  fmt.Sprintf("Missing ladok permission id: %d (%s), permission level: %q", permissionID, c.translateID(ctx, permissionID), c.translatePermission(ctx, data["my"])),
				Type: "Ladok permission",
			})
			continue
		}

		if !c.comparePermission(ctx, ladokPermission, myPermission) {
			// ladokPermission does not reach myPermission
			myPermission := data["my"]
			internalError = append(internalError, ladoktypes.InternalError{
				Msg:  fmt.Sprintf("Not sufficient permission: %q for id: %d (%s)", c.translatePermission(ctx, myPermission), permissionID, c.translateID(ctx, permissionID)),
				Type: "Ladok permission",
			})
		}
	}
	if len(internalError) > 0 {
		e.Internal = internalError
		return e

	}
	return nil
}

// comparePermission compare l with m permission.
func (c *Client) comparePermission(ctx context.Context, l, m int64) bool {
	ctx, span := c.tp.Start(ctx, "goladok3.comparePermission")
	defer span.End()

	if l == m {
		return true
	}

	switch l {
	case 4:
		if m == 6 {
			return true
		}
	case 6:
		if m == 4 {
			return false
		}
	}

	return false
}

// permissionUnify convert ladok permission structure to something that's easier to compare.
func (c *Client) permissionUnify(ctx context.Context, l ladoktypes.KataloginformationBehorighetsprofil, p Permissions) (permissions map[int64]map[string]int64, err error) {
	ctx, span := c.tp.Start(ctx, "goladok3.permissionUnify")
	defer span.End()

	if len(l.Systemaktiviteter) == 0 {
		return nil, ErrNotSufficientPermissions
	}
	if len(p) == 0 {
		return nil, ErrNoPermissionProvided
	}

	permissions = make(map[int64]map[string]int64)

	parse := func(ctx context.Context, permission, className string, classMap map[string]int64, key int64, store map[int64]map[string]int64) {
		ctx, span := c.tp.Start(ctx, "goladok3.parse")
		defer span.End()

		switch permission {
		case "rattighetsniva.las":
			classMap[className] = 4
		case "rattighetsniva.lokal":
			classMap[className] = 6
		default:
			classMap[className] = 0
		}
	}

	for key, permission := range p {
		classMyMap := make(map[string]int64)
		parse(ctx, permission, "my", classMyMap, key, permissions)
		permissions[key] = classMyMap
	}

	for _, sys := range l.Systemaktiviteter {
		classMap := make(map[string]int64)
		key := sys.ID
		permission := sys.Rattighetsniva
		parse(ctx, permission, "ladok", classMap, key, permissions)
		val, ok := permissions[key]
		if ok {
			val["ladok"] = classMap["ladok"]
			permissions[key] = val
		} else {
			permissions[key] = classMap
		}
	}

	return permissions, nil
}

func (c *Client) translatePermission(ctx context.Context, p int64) string {
	ctx, span := c.tp.Start(ctx, "goladok3.translatePermission")
	defer span.End()

	switch p {
	case 0:
		return "NoPermission"
	case 4:
		return "rattighetsniva.las"
	case 6:
		return "rattighetsniva.lokal"
	default:
		return "Undefined"
	}
}

func (c *Client) translateID(ctx context.Context, p int64) string {
	ctx, span := c.tp.Start(ctx, "goladok3.translateID")
	defer span.End()

	switch p {
	case 90019:
		return "uppfoljning.feeds"
	case 51001:
		return "studiedeltagande.las"
	case 61001:
		return "studentinformation.lasa"
	case 11004:
		return "kataloginformation.las"
	case 860131:
		return "extintegration.lasa"
	default:
		return "Undefined"
	}
}

func (c *Client) environment(ctx context.Context) (string, error) {
	ctx, span := c.tp.Start(ctx, "goladok3.environment")
	defer span.End()

	switch c.certificate.Subject.OrganizationalUnit[1] {
	case ladoktypes.EnvIntTestAPI:
		return ladoktypes.EnvIntTestAPI, nil
	case ladoktypes.EnvProdAPI:
		return ladoktypes.EnvProdAPI, nil
	case ladoktypes.EnvTestAPI:
		return ladoktypes.EnvTestAPI, nil
	default:
		return "", ErrNoEnvFound
	}
}
