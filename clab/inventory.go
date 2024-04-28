package clab

import (
    "os"
	_ "embed"
    "text/template"
    "github.com/srl-labs/containerlab/types"
)

// Embedding both inventory templates
//go:embed inventory_ansible.go.tpl
var ansibleInvT string

//go:embed inventory_nornir.go.tpl
var nornirInvT string

// General inventory node structure, usable by both inventory types
type InventoryNode struct {
    *types.NodeConfig
}

// Ansible-specific inventory structures
type KindProps struct {
    Username    string
    Password    string
    NetworkOS   string
    AnsibleConn string
}

type AnsibleInventory struct {
    Kinds  map[string]*KindProps
    Nodes  map[string][]*InventoryNode
    Groups map[string][]*InventoryNode
}

// Nornir-specific inventory structure
type NornirInventory struct {
    Nodes []*InventoryNode
}

// CLab modifications to include both inventory generation methods
type CLab struct {
    // Existing CLab fields
    Nodes map[string]*types.Node
    TopoPaths *TopologyPaths
}

type TopologyPaths struct {
    AnsibleInventoryFileAbsPath string
    NornirInventoryFileAbsPath string
}

// Generate inventories for both Ansible and Nornir
func (c *CLab) GenerateInventories() error {
    // Generate Ansible inventory
    if err := c.generateAnsibleInventory(); err != nil {
        return err
    }

    // Generate Nornir inventory
    return c.generateNornirInventory()
}

func (c *CLab) generateAnsibleInventory() error {
    fPath := c.TopoPaths.AnsibleInventoryFileAbsPath()
    f, err := os.Create(fPath)
    if err != nil {
        return err
    }
    defer f.Close()

    inv := AnsibleInventory{
        Kinds:  make(map[string]*KindProps),
        Nodes:  make(map[string][]*InventoryNode),
        Groups: make(map[string][]*InventoryNode),
    }

    // Logic to populate inv from c.Nodes

    t, err := template.New("ansible").Parse(ansibleInvT)
    if err != nil {
        return err
    }
    return t.Execute(f, inv)
}

func (c *CLab) generateNornirInventory() error {
	fPath := c.TopoPaths.NornirInventoryHostsFileAbsPath()
    f, err := os.Create(fPath)
    if err != nil {
        return err
    }
    defer f.Close()

    inv := NornirInventory{
        Nodes: make([]*InventoryNode, 0),
    }

    // Logic to populate inv from c.Nodes

    t, err := template.New("nornir").Parse(nornirInvT)
    if err != nil {
        return err
    }
    return t.Execute(f, inv)
}
