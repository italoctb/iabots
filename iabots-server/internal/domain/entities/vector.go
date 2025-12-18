package entities

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Vector []float32

func NewVector(values []float32) Vector {
	return Vector(values)
}

func (v *Vector) Scan(value interface{}) error {
	if value == nil {
		*v = nil
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("unable to scan Vector from: %v", value)
	}

	str = strings.TrimPrefix(str, "{")
	str = strings.TrimSuffix(str, "}")
	parts := strings.Split(str, ",")

	vec := make(Vector, len(parts))
	for i, part := range parts {
		val, err := strconv.ParseFloat(part, 32)
		if err != nil {
			return err
		}
		vec[i] = float32(val)
	}

	*v = vec
	return nil
}

func (v Vector) Value() (driver.Value, error) {
	if v == nil {
		return nil, nil
	}

	parts := make([]string, len(v))
	for i, val := range v {
		parts[i] = strconv.FormatFloat(float64(val), 'f', -1, 32)
	}
	return "{" + strings.Join(parts, ",") + "}", nil
}

func (v Vector) ToString() string {
	if len(v) == 0 {
		return "ARRAY[]::double precision[]"
	}
	builder := strings.Builder{}
	builder.WriteString("ARRAY[")
	for i, val := range v {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(strconv.FormatFloat(float64(val), 'f', -1, 32))
	}
	builder.WriteString("]::double precision[]")
	return builder.String()
}

/**
CREATE OR REPLACE FUNCTION cosine_similarity(a double precision[], b double precision[])
RETURNS double precision AS $$
DECLARE
    dot_product double precision := 0;
    norm_a double precision := 0;
    norm_b double precision := 0;
    i integer;
BEGIN
    IF array_length(a, 1) <> array_length(b, 1) THEN
        RAISE EXCEPTION 'The input arrays must have the same length';
    END IF;

    FOR i IN 1..array_length(a, 1)
    LOOP
        dot_product := dot_product + (a[i] * b[i]);
        norm_a := norm_a + (a[i] * a[i]);
        norm_b := norm_b + (b[i] * b[i]);
    END LOOP;

    IF norm_a = 0 OR norm_b = 0 THEN
        RETURN 0;
    ELSE
        RETURN dot_product / (sqrt(norm_a) * sqrt(norm_b));
    END IF;
END;
$$ LANGUAGE plpgsql;
**/
