package handlers

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
        v.err = errors.New("El parametro no puede ser nulo.")
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
        return "", errors.New("Nombre invalido.")
    } else {
        return "", nil
    }
}

func (v *Validation) AsLogin() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^[a-zA-Z_]+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Login invalido.")
    } else {
        return "", nil
    }
}

func (v *Validation) AsPassword() (string, error) {
    if v.err != nil {
        return "", v.err
    }
    re := regexp.MustCompile(`^\w+$`)
    if len(v.parameter) > 0 && re.MatchString(v.parameter) {
        return v.parameter, nil
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return "", errors.New("Password invalida.")
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
        return "", errors.New("Texto invalido.")
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
            return 0, errors.New("Precio invalido.")
        } else {
            return f, nil
        }
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return 0, errors.New("Precio invalido.")
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
            return 0, errors.New("Fecha invalida.")
        } else {
            return u, nil
        }
    } else if len(v.parameter) > 0 && !re.MatchString(v.parameter) {
        return 0, errors.New("Fecha invalida.")
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
        return "", errors.New("Numero de patente invalido.")
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
        return "", errors.New("Numero de chasis invalido.")
    } else {
        return "", nil
    }
}
