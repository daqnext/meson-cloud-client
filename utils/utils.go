package utils

import (
    "os"
    "path/filepath"

    "daqnext/meson-cloud-client/logger"
)

func RelPathAndCheck(DIR_PREFIX, path string) (string, bool, error) {
    //check abs path or relative path
    is_abs := filepath.IsAbs(path)
    if is_abs {
        exist, err := PathExist(path)
        if err != nil {
            return "", false, err
        }
        return path, exist, nil
    }

    logger.L.Debugln("new path", path)

    // workding dir
    path1, exist, err := toRelPath(path)
    if err != nil {
        return "", false, err
    }

    if exist {
        path = path1
        return path, exist, nil
    }

    // binary dir
    path = filepath.Join(DIR_PREFIX, path)
    path, exist, err = toRelPath(path)
    if err != nil {
        return "", false, err
    }

    return path, exist, nil
}

func toRelPath(path string) (string, bool, error) {
    // to rel directory
    path, err := filepath.Abs(path)
    if err != nil {
        return "", false, err
    }
    exist, err := PathExist(path)
    if err != nil {
        return "", false, err
    }
    return path, exist, nil
}

func PathExist(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
