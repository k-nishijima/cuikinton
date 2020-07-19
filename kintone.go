package cuikinton

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/kintone/go-kintone"
)

type KintoneConfig struct {
	Domain       string
	User         string
	Password     string
	ApiToken     string
	AppId        uint64
	GuestSpaceId uint64
}

func getKintoneApp(c KintoneConfig) *kintone.App {
	return &kintone.App{
		Domain:       c.Domain,
		User:         c.User,
		Password:     c.Password,
		ApiToken:     c.ApiToken,
		AppId:        c.AppId,
		GuestSpaceId: c.GuestSpaceId,
	}
}

func getRecords(app *kintone.App) ([]*kintone.Record, error) {
	if result, err := app.GetRecords(nil, "order by $id desc"); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func getHeadline(r *kintone.Record) string {
	txt := ""

	keys := make([]string, len(r.Fields), len(r.Fields))
	i := 0
	for key := range r.Fields {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		a := keys[i]
		b := keys[j]
		return a < b
	})

	for _, k := range keys {
		col := kintoneField2String(r.Fields[k])
		txt += col + " | "
	}

	return txt
}

func prettyPrintKintoneRecord(r *kintone.Record) string {
	recJSON, _ := json.MarshalIndent(r, "", "  ")
	return string(recJSON)
}

func kintoneField2String(f interface{}) string {
	delimiter := ","
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	switch f.(type) {
	case kintone.SingleLineTextField:
		singleLineTextField := f.(kintone.SingleLineTextField)
		return string(singleLineTextField)
	case kintone.MultiLineTextField:
		multiLineTextField := f.(kintone.MultiLineTextField)
		return string(multiLineTextField)
	case kintone.RichTextField:
		richTextField := f.(kintone.RichTextField)
		return string(richTextField)
	case kintone.DecimalField:
		decimalField := f.(kintone.DecimalField)
		return string(decimalField)
	case kintone.CalcField:
		calcField := f.(kintone.CalcField)
		return string(calcField)
	case kintone.RadioButtonField:
		radioButtonField := f.(kintone.RadioButtonField)
		return string(radioButtonField)
	case kintone.LinkField:
		linkField := f.(kintone.LinkField)
		return string(linkField)
	case kintone.StatusField:
		statusField := f.(kintone.StatusField)
		return string(statusField)
	case kintone.RecordNumberField:
		recordNumberField := f.(kintone.RecordNumberField)
		return string(recordNumberField)
	case kintone.CheckBoxField:
		checkBoxField := f.(kintone.CheckBoxField)
		return strings.Join(checkBoxField, delimiter)
	case kintone.MultiSelectField:
		multiSelectField := f.(kintone.MultiSelectField)
		return strings.Join(multiSelectField, delimiter)
	case kintone.CategoryField:
		categoryField := f.(kintone.CategoryField)
		return strings.Join(categoryField, delimiter)
	case kintone.SingleSelectField:
		singleSelect := f.(kintone.SingleSelectField)
		return singleSelect.String
	case kintone.FileField:
		fileField := f.(kintone.FileField)
		files := make([]string, 0, len(fileField))
		for _, file := range fileField {
			files = append(files, file.Name)
		}
		return strings.Join(files, delimiter)
	case kintone.DateField:
		dateField := f.(kintone.DateField)
		if dateField.Valid {
			// UTCをJSTに変換
			nowJST := dateField.Date.UTC().In(jst)
			return nowJST.Format("2006-01-02")
		} else {
			return ""
		}
	case kintone.TimeField:
		timeField := f.(kintone.TimeField)
		if timeField.Valid {
			// UTCをJSTに変換
			nowJST := timeField.Time.UTC().In(jst)
			return nowJST.Format("15:04:05")
		} else {
			return ""
		}
	case kintone.DateTimeField:
		dateTimeField := f.(kintone.DateTimeField)
		if dateTimeField.Valid {
			// UTCをJSTに変換
			nowJST := dateTimeField.Time.UTC().In(jst)
			return nowJST.Format(time.RFC3339)
		} else {
			return ""
		}
	case kintone.UserField:
		userField := f.(kintone.UserField)
		users := make([]string, 0, len(userField))
		for _, user := range userField {
			users = append(users, user.Code)
		}
		return strings.Join(users, delimiter)
	case kintone.OrganizationField:
		organizationField := f.(kintone.OrganizationField)
		organizations := make([]string, 0, len(organizationField))
		for _, organization := range organizationField {
			organizations = append(organizations, organization.Code)
		}
		return strings.Join(organizations, delimiter)
	case kintone.GroupField:
		groupField := f.(kintone.GroupField)
		groups := make([]string, 0, len(groupField))
		for _, group := range groupField {
			groups = append(groups, group.Code)
		}
		return strings.Join(groups, delimiter)
	case kintone.AssigneeField:
		assigneeField := f.(kintone.AssigneeField)
		users := make([]string, 0, len(assigneeField))
		for _, user := range assigneeField {
			users = append(users, user.Code)
		}
		return strings.Join(users, delimiter)
	case kintone.CreatorField:
		creatorField := f.(kintone.CreatorField)
		return creatorField.Code
	case kintone.ModifierField:
		modifierField := f.(kintone.ModifierField)
		return modifierField.Code
	case kintone.CreationTimeField:
		creationTimeField := f.(kintone.CreationTimeField)
		return time.Time(creationTimeField).Format(time.RFC3339)
	case kintone.ModificationTimeField:
		modificationTimeField := f.(kintone.ModificationTimeField)
		return time.Time(modificationTimeField).Format(time.RFC3339)
	case kintone.SubTableField:
		return "" // unsupported
	}
	return ""
}
