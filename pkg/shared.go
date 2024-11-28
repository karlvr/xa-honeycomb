package pkg

import "github.com/expr-lang/expr/vm"

type Command struct {
	Command_str string `yaml:"command_str,omitempty" json:"command_str,omitempty"`
	Command     interface{}
}

type Dataref struct {
	Dataref_str string `yaml:"dataref_str,omitempty" json:"dataref_str,omitempty"`
	Dataref     interface{}
	Operator    string  `yaml:"operator,omitempty" json:"operator,omitempty"`
	Threshold   float32 `yaml:"threshold,omitempty" json:"threshold,omitempty"`
	Index       int     `yaml:"index,omitempty" json:"index,omitempty"`
	Expr        *vm.Program
	Env         map[string]interface{}
}
type Data struct {
	Dataref_str string `yaml:"dataref_str,omitempty" json:"dataref_str,omitempty"`
	Dataref     interface{}
}

type BravoProfile struct {
	ProfileType string    `yaml:"profile_type,omitempty" json:"profile_type,omitempty"`
	Condition   string    `yaml:"condition,omitempty" json:"condition,omitempty"`
	Datarefs    []Dataref `yaml:"datarefs,omitempty" json:"datarefs,omitempty"`
	Commands    []Command `yaml:"commands,omitempty" json:"commands,omitempty"`
	Data        []Data    `yaml:"data,omitempty" json:"data,omitempty"`
	On          func()    `json:"-"`
	Off         func()    `json:"-"`
}

type Profile struct {
	Name string `yaml:"name" json:"name"`
	// AP Knobs
	AP_HDG BravoProfile `yaml:"ap_hdg,omitempty" json:"ap_hdg,omitempty"`
	AP_VS  BravoProfile `yaml:"ap_vs,omitempty" json:"ap_vs,omitempty"`
	AP_ALT BravoProfile `yaml:"ap_alt,omitempty" json:"ap_alt,omitempty"`
	AP_IAS BravoProfile `yaml:"ap_ias,omitempty" json:"ap_ias,omitempty"`
	AP_CRS BravoProfile `yaml:"ap_crs,omitempty" json:"ap_crs,omitempty"`
	// LEDs
	BUS_VOLTAGE        BravoProfile `yaml:"bus_voltage,omitempty" json:"bus_voltage,omitempty"`
	HDG                BravoProfile `yaml:"hdg,omitempty" json:"hdg,omitempty"`
	NAV                BravoProfile `yaml:"nav,omitempty" json:"nav,omitempty"`
	ALT                BravoProfile `yaml:"alt,omitempty" json:"alt,omitempty"`
	APR                BravoProfile `yaml:"apr,omitempty" json:"apr,omitempty"`
	VS                 BravoProfile `yaml:"vs,omitempty" json:"vs,omitempty"`
	AP                 BravoProfile `yaml:"ap,omitempty" json:"ap,omitempty"`
	IAS                BravoProfile `yaml:"ias,omitempty" json:"ias,omitempty"`
	REV                BravoProfile `yaml:"rev,omitempty" json:"rev,omitempty"`
	AP_STATE           BravoProfile `yaml:"ap_state,omitempty" json:"ap_state,omitempty"`
	GEAR               BravoProfile `yaml:"gear,omitempty" json:"gear,omitempty"`
	RETRACTABLE_GEAR   BravoProfile `yaml:"retractable_gear,omitempty" json:"retractable_gear,omitempty"`
	MASTER_WARN        BravoProfile `yaml:"master_warn,omitempty" json:"master_warn,omitempty"`
	MASTER_CAUTION     BravoProfile `yaml:"master_caution,omitempty" json:"master_caution,omitempty"`
	FIRE               BravoProfile `yaml:"fire,omitempty" json:"fire,omitempty"`
	OIL_LOW_PRESSURE   BravoProfile `yaml:"oil_low_pressure,omitempty" json:"oil_low_pressure,omitempty"`
	FUEL_LOW_PRESSURE  BravoProfile `yaml:"fuel_low_pressure,omitempty" json:"fuel_low_pressure,omitempty"`
	ANTI_ICE           BravoProfile `yaml:"anti_ice,omitempty" json:"anti_ice,omitempty"`
	ENG_STARTER        BravoProfile `yaml:"eng_starter,omitempty" json:"eng_starter,omitempty"`
	APU                BravoProfile `yaml:"apu,omitempty" json:"apu,omitempty"`
	VACUUM             BravoProfile `yaml:"vacuum,omitempty" json:"vacuum,omitempty"`
	HYDRO_LOW_PRESSURE BravoProfile `yaml:"hydro_low_pressure,omitempty" json:"hydro_low_pressure,omitempty"`
	AUX_FUEL_PUMP      BravoProfile `yaml:"aux_fuel_pump,omitempty" json:"aux_fuel_pump,omitempty"`
	PARKING_BRAKE      BravoProfile `yaml:"parking_brake,omitempty" json:"parking_brake,omitempty"`
	VOLT_LOW           BravoProfile `yaml:"volt_low,omitempty" json:"volt_low,omitempty"`
	DOORS              BravoProfile `yaml:"doors,omitempty" json:"doors,omitempty"`
}