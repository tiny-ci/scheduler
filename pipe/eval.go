package pipe

import (
	"errors"
    "regexp"
	"strings"
)

const SemverRegex =
    `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)` +
    `(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?` +
    `(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

func isSemver(name string) bool {
    _, err := regexp.MatchString(SemverRegex, name)
    return err == nil
}

func startsWith(name string, prefix string) bool {
    return strings.HasPrefix(name, prefix)
}

func endsWith(name string, suffix string) bool {
    return strings.HasSuffix(name, suffix)
}

func EvalExpr(expr string, name string) (bool, error) {
    lcExpr := strings.ToLower(expr)

    if strings.HasPrefix(lcExpr, "\\semver") {
        return isSemver(name), nil
    }

    if strings.HasPrefix(lcExpr, "\\startswith ") {
        return startsWith(name, expr[12:]), nil
    }

    if strings.HasPrefix(lcExpr, "\\endswith ") {
        return endsWith(name, expr[11:]), nil
    }

    return false, errors.New("unknown expression ")
}
