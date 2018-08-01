package main

import (
	"encoding/json"
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/log"
	"github.com/google/lvmd/parser"
	"os/exec"
	"strconv"
	"strings"
)

const separator = "<:SEP:>"

func main() {
	//[]*parser.LV
	//lvs, err := ListLV("vg-k8s")
	vg, err := GetVg("centos")
	if err != nil {
		log.Error(err)
		return
	}
	lvsBytes, err := json.Marshal(vg)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Print(string(lvsBytes))
}

// ListLV lists lvm volumes
func ListLV(listspec string) ([]*parser.LV, error) {
	cmd := exec.Command("lvs", "--units=b", "--separator=<:SEP:>", "--nosuffix", "--noheadings",
		"-o", "lv_name,lv_size,lv_uuid,lv_attr,copy_percent,lv_kernel_major,lv_kernel_minor,lv_tags", "--nameprefixes", "-a", listspec)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	outLines := strings.Split(outStr, "\n")
	lvs := make([]*parser.LV, len(outLines))
	for i, line := range outLines {
		line = strings.TrimSpace(line)
		lv, err := parser.ParseLV(line)
		if err != nil {
			return nil, err
		}
		lvs[i] = lv
	}
	return lvs, nil
}

// ListLV lists lvm volumes
func GetVg(spec string) (*VG, error) {
	cmd := exec.Command("vgs", "--units=b", "--separator=<:SEP:>", "--nosuffix", "--noheadings", "-o", "vg_name,vg_uuid,vg_size,vg_free,vg_attr,vg_tags", "--nameprefixes", spec)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	vg, err := ParseVG(outStr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return vg, nil
}

type VG struct {
	Name string
	UUID string
	Size uint64
	Free uint64
	Tags []string
}

// ParseLV parses a line from lvs
func ParseVG(line string) (*VG, error) {
	components := strings.Split(line, separator)
	if len(components) != 6 {
		return nil, fmt.Errorf("expected 8 components, got %d", len(components))
	}

	fields := map[string]string{}
	for _, c := range components {
		idx := strings.Index(c, "=")
		if idx == -1 {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		key := c[0:idx]
		value := c[idx+1 : len(c)]
		if len(value) < 2 {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		if value[0] != '\'' || value[len(value)-1] != '\'' {
			return nil, fmt.Errorf("failed to parse component '%s'", c)
		}
		value = value[1 : len(value)-1]
		fields[key] = value
	}

	size, err := strconv.ParseUint(fields["LVM2_VG_SIZE"], 10, 64)
	if err != nil {
		return nil, err
	}

	free, err := strconv.ParseUint(fields["LVM2_VG_FREE"], 10, 64)
	if err != nil {
		return nil, err
	}

	return &VG{
		Name: fields["LVM2_VG_NAME"],
		UUID: fields["LVM2_VG_UUID"],
		Size: size,
		Free: free,
		Tags: strings.Split(fields["LVM2_VG_TAGS"], ","),
	}, nil
}
