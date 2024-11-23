//go:build !test

package xplane

//go:generate mockgen -destination=./__mocks__/xplane.go -package=mocks -source=xplane.go

import (
	"github.com/xairline/goplane/extra"
	"github.com/xairline/goplane/xplm/plugins"
	"github.com/xairline/xa-honeycomb/pkg"
	"github.com/xairline/xa-honeycomb/pkg/honeycomb"
	"runtime"
	"sync"
)

var VERSION = "development"

type XplaneService interface {
	// init
	onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin)
	onPluginStart()
	onPluginStop()
	// flight loop
	flightLoop(elapsedSinceLastCall, elapsedTimeSinceLastFlightLoop float32, counter int, ref interface{}) float32
}

type xplaneService struct {
	Plugin       *extra.XPlanePlugin
	BravoService honeycomb.BravoService
	Logger       pkg.Logger
}

var xplaneSvcLock = &sync.Mutex{}
var xplaneSvc XplaneService

func NewXplaneService(
	logger pkg.Logger,
) XplaneService {
	if xplaneSvc != nil {
		logger.Info("Xplane SVC has been initialized already")
		return xplaneSvc
	} else {
		logger.Info("Xplane SVC: initializing")
		xplaneSvcLock.Lock()
		defer xplaneSvcLock.Unlock()
		xplaneSvc := &xplaneService{
			Plugin:       extra.NewPlugin("X Airline Snow - "+VERSION, "com.github.xairline.xa-snow", "show accumulated snow in X-Plane's world"),
			BravoService: honeycomb.NewBravoService(logger, "TODO"),
		}
		xplaneSvc.Plugin.SetPluginStateCallback(xplaneSvc.onPluginStateChanged)
		xplaneSvc.Plugin.SetMessageHandler(xplaneSvc.messageHandler)
		return xplaneSvc
	}
}

func (s *xplaneService) onPluginStateChanged(state extra.PluginState, plugin *extra.XPlanePlugin) {
	switch state {
	case extra.PluginStart:
		s.onPluginStart()
	case extra.PluginStop:
		s.onPluginStop()
	case extra.PluginEnable:
		s.Logger.Infof("Plugin: %s enabled", plugin.GetName())
	case extra.PluginDisable:
		s.Logger.Infof("Plugin: %s disabled", plugin.GetName())
	}
}

func (s *xplaneService) onPluginStart() {
	s.Logger.Info("Plugin started")
	runtime.GOMAXPROCS(runtime.NumCPU())

	//systemPath := utilities.GetSystemPath()
	//err := godotenv.Load(s.configFilePath)
	//if err != nil {
	//	s.Logger.Errorf("Some error occured. Err: %s", err)
	//}
	//if os.Getenv("OVERRIDE") == "true" {
	//	s.override = true
	//}
	//
	//// API drefs are available at plugin start
	//s.lat_dr, _ = dataAccess.FindDataRef("sim/flightmodel/position/latitude")
	//s.lon_dr, _ = dataAccess.FindDataRef("sim/flightmodel/position/longitude")
	//s.weatherMode_dr, _ = dataAccess.FindDataRef("sim/weather/region/weather_source")
	//s.sysTime_dr, _ = dataAccess.FindDataRef("sim/time/use_system_time")
	//s.simCurrentMonth_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/current_month")
	//s.simCurrentDay_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/current_day")
	//s.simLocalHours_dr, _ = dataAccess.FindDataRef("sim/cockpit2/clock_timer/local_time_hours")
	//
	//// start with delay to let the dust settle
	//processing.RegisterFlightLoopCallback(s.flightLoop, 5.0, nil)
	//
	//// setup menu
	//menuId := menus.FindPluginsMenu()
	//menuContainerId := menus.AppendMenuItem(menuId, "X Airline Snow", 0, false)
	//s.myMenuId = menus.CreateMenu("X Airline Snow", menuId, menuContainerId, s.menuHandler, nil)
	//s.myMenuItemIndex = menus.AppendMenuItem(s.myMenuId, "Toggle Override", 0, true)
	//if s.override {
	//	menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	//} else {
	//	menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	//}
	//
	//// set internal vars to known "no snow" state
	//s.snowNow, s.rwySnowCover, s.iceNow = s.p2x.SnowDepthToXplaneSnowNow(0)
}

func (s *xplaneService) onPluginStop() {
	s.Logger.Info("Plugin stopped")
}

// flightloop, high freq code!
func (s *xplaneService) flightLoop(
	elapsedSinceLastCall,
	elapsedTimeSinceLastFlightLoop float32,
	counter int,
	ref interface{},
) float32 {

	// flightloop start is the first point in time where the time datarefs are set correctly
	//if s.loopCnt == 0 {
	//	s.loopCnt++
	//	s.Logger.Info("Flightloop starting, kicking off")
	//
	//	// delayed init
	//	if !initDrefs(s) {
	//		return 0 // Bye, if we don't have them by now we will never get them
	//	}
	//
	//	sys_time := dataAccess.GetIntData(s.sysTime_dr) == 1
	//	day := dataAccess.GetIntData(s.simCurrentDay_dr)
	//	month := dataAccess.GetIntData(s.simCurrentMonth_dr)
	//	hour := dataAccess.GetIntData(s.simLocalHours_dr)
	//
	//	go func() {
	//		for {
	//			err := honeycomb.bravoSvc.DownloadAndProcessBravoFile(sys_time, month, day, hour)
	//			if err != nil {
	//				s.Logger.Errorf("Download bravo file failed: %v", err)
	//			} else {
	//				// TODO: was this looping forever?
	//				break
	//			}
	//			// TODO: disabled - auto NOAA update
	//		}
	//	}()
	//
	//	return 10.0
	//}
	//
	//if !s.override {
	//	weatherMode := dataAccess.GetIntData(s.weatherMode_dr)
	//	if weatherMode != 1 {
	//		// weather mode is not RW, we don't do anything to avoid snow on people's summer view
	//		return 5.0
	//	}
	//}
	//
	//if !s.BravoService.IsReady() {
	//	return 2.0
	//}
	//
	//// throttle update computations
	//s.loopCnt++
	//if s.loopCnt%8 == 0 {
	//	lat := dataAccess.GetFloatData(s.lat_dr)
	//	lon := dataAccess.GetFloatData(s.lon_dr)
	//	snowDepth_n := s.BravoService.GetSnowDepth(lat, lon)
	//
	//	// some exponential smoothing
	//	const alpha = float32(0.7)
	//	s.snowDepth = alpha*snowDepth_n + (1-alpha)*s.snowDepth
	//
	//	// If we have no accumulated snow leave the datarefs alone and let X-Plane do its weather effects
	//	if s.snowDepth < 0.001 && !s.override {
	//		return -1
	//	}
	//
	//	s.snowNow, s.rwySnowCover, s.iceNow = s.p2x.SnowDepthToXplaneSnowNow(s.snowDepth)
	//}
	//
	//// If we have no accumulated snow leave the datarefs alone and let X-Plane do its weather effects
	//if s.snowDepth < 0.001 && !s.override {
	//	return -1
	//}
	//
	//dataAccess.SetFloatData(s.snow_dr, s.snowNow)
	//dataAccess.SetFloatData(s.rwySnowCover_dr, s.rwySnowCover)
	//dataAccess.SetFloatData(s.ice_dr, s.iceNow)

	return -1
}

func (s *xplaneService) messageHandler(message plugins.Message) {
	if message.MessageId == plugins.MSG_PLANE_LOADED || message.MessageId == plugins.MSG_SCENERY_LOADED {
		s.Logger.Info("Plane/Scenery loaded")
	}
}

func (s *xplaneService) menuHandler(menuRef interface{}, itemRef interface{}) {
	//s.override = !s.override
	//
	//if s.override {
	//	menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Checked)
	//} else {
	//	menus.CheckMenuItem(s.myMenuId, s.myMenuItemIndex, menus.Menu_Unchecked)
	//}
	//
	//// write to config
	//err := godotenv.Write(map[string]string{
	//	"OVERRIDE": strconv.FormatBool(s.override),
	//}, s.configFilePath)
	//if err != nil {
	//	s.Logger.Errorf("Error writing to config: %v", err)
	//}

	s.Logger.Infof("menu clicked: %v", itemRef)
}
