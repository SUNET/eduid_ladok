package httpserver

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

func (s *Service) bindRequest(c *gin.Context, v any) error {
	if c.ContentType() == gin.MIMEJSON {
		if err := c.ShouldBindJSON(v); err != nil {
			s.logger.Warn("bind_json_error", "error", err)
		}
	}
	if err := s.bindRequestQuery(c, v); err != nil {
		s.logger.Warn("bind_query_map_error", "error", err)
	}
	if err := c.ShouldBindQuery(v); err != nil {
		s.logger.Warn("bind_query_error", "error", err)
	}
	return c.ShouldBindUri(v)
}

func (s *Service) bindRequestQuery(c *gin.Context, v any) error {
	refV := reflect.ValueOf(v).Elem()
	refT := reflect.ValueOf(v).Elem().Type()
	for field := range refT.Fields() {
		field := field
		fieldType := field.Type
		fieldKey := field.Tag.Get("form")
		if fieldKey == "" {
			fieldKey = field.Name
		}
		switch fieldType.String() {
		case "map[string]string":
			v := c.QueryMap(fieldKey)
			if len(v) == 0 {
				continue
			}
			refV.FieldByName(field.Name).Set(reflect.ValueOf(v))
		case "*map[string]string":
			v := c.QueryMap(fieldKey)
			if len(v) == 0 {
				continue
			}
			refV.FieldByName(field.Name).Set(reflect.ValueOf(&v))
		case "map[string][]string":
			v := make(map[string][]string)
			for key := range c.QueryMap(fieldKey) {
				v[key] = c.QueryArray(fmt.Sprintf("%s[%s]", fieldKey, key))
			}
			if len(v) == 0 {
				continue
			}
			refV.FieldByName(field.Name).Set(reflect.ValueOf(v))
		case "*map[string][]string":
			v := make(map[string][]string)
			for key := range c.QueryMap(fieldKey) {
				v[key] = c.QueryArray(fmt.Sprintf("%s[%s]", fieldKey, key))
			}
			if len(v) == 0 {
				continue
			}
			refV.FieldByName(field.Name).Set(reflect.ValueOf(&v))
		}
	}
	return nil
}
