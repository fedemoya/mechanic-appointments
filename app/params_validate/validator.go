package params_validate

import (
    "strconv"
    "errors"
    "regexp"
)

type Validation struct {
    parameter string
    err error
}

func Validate(parameter string) *Validation {
    return &Validation{parameter: parameter, err: nil}
}

func (v *Validation) Required() *Validation {
    if len(v.parameter) == 0 {
        v.err = errors.New("The parameter cannot be null.")
    }
    return v
}

func (v *Validation) AsName() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Invalid Name.")
    } else {
        return "", nil
    }
}

func (v *Validation) AsText() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^[\w\s]+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Invalid Text.")
    } else {
        return "", nil
    }
}

func (v *Validation) AsPrice() (float64, error) {
    if v.err != nil {
        return 0, v.err
    }
    re := regexp.MustCompile(`^\d+(\.\d+\d+)?$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        f, err := strconv.ParseFloat(v.parameter, 64)
        if err != nil {
            return 0, errors.New("Invalid Price.")
        } else {
            return f, nil
        }
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return 0, errors.New("Invalid Price.")
    } else {
        return 0, nil
    }
}

func (v *Validation) AsEpochTime() (uint64, error) {
    if v.err != nil {
        return 0, v.err
    }
    re := regexp.MustCompile(`^\d+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        u, err := strconv.ParseUint(v.parameter, 10, 64)
        if err != nil {
            return 0, errors.New("Invalid Epoch Time.")
        } else {
            return u, nil
        }
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return 0, errors.New("Invalid Epoch Time.")
    } else {
        return 0, nil
    }
}

func (v *Validation) AsPlateNumber() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^[A-Z0-9]+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Invalid Plate Number.")
    } else {
        return "", nil
    }
}

func (v *Validation) AsChassisNumber() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^[A-Z0-9]+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Invalid Chassis Number.")
    } else {
        return "", nil
    }
}
