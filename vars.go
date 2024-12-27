package aichat

import (
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const (
	userAgent       = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0"
	contentTypeJson = "application/json"
)

var (
	Json            = jsoniter.ConfigCompatibleWithStandardLibrary
	hostURL, _      = url.Parse("https://chatgpt.com")
	cores           = []int{16, 24, 32}
	screens         = []int{3000, 4000, 6000}
	timeLocation, _ = time.LoadLocation("Asia/Shanghai")
	timeLayout      = "Mon Jan 2 2006 15:04:05"
	dataBuildId     = "prod-0749372c57dc2f5a1c6aa5f4f9802a682aebf139"
	specialCases    = map[string]string{
		"window.Math":            "[object Math]",
		"window.Reflect":         "[object Reflect]",
		"window.performance":     "[object Performance]",
		"window.localStorage":    "[object Storage]",
		"window.Object":          "function Object() { [native code] }",
		"window.Reflect.set":     "function set() { [native code] }",
		"window.performance.now": "function () { [native code] }",
		"window.Object.create":   "function create() { [native code] }",
		"window.Object.keys":     "function keys() { [native code] }",
		"window.Math.random":     "function random() { [native code] }",
	}
	navigatorMap = []string{
		"registerProtocolHandler−function registerProtocolHandler() { [native code] }",
		"storage−[object StorageManager]",
		"locks−[object LockManager]",
		"appCodeName−Mozilla",
		"permissions−[object Permissions]",
		"appVersion−5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"share−function share() { [native code] }",
		"webdriver−false",
		"managed−[object NavigatorManagedData]",
		"canShare−function canShare() { [native code] }",
		"vendor−Google Inc.",
		"vendor−Google Inc.",
		"mediaDevices−[object MediaDevices]",
		"vibrate−function vibrate() { [native code] }",
		"storageBuckets−[object StorageBucketManager]",
		"mediaCapabilities−[object MediaCapabilities]",
		"getGamepads−function getGamepads() { [native code] }",
		"bluetooth−[object Bluetooth]",
		"share−function share() { [native code] }",
		"cookieEnabled−true",
		"virtualKeyboard−[object VirtualKeyboard]",
		"product−Gecko",
		"mediaDevices−[object MediaDevices]",
		"canShare−function canShare() { [native code] }",
		"getGamepads−function getGamepads() { [native code] }",
		"product−Gecko",
		"xr−[object XRSystem]",
		"clipboard−[object Clipboard]",
		"storageBuckets−[object StorageBucketManager]",
		"unregisterProtocolHandler−function unregisterProtocolHandler() { [native code] }",
		"productSub−20030107",
		"login−[object NavigatorLogin]",
		"vendorSub−",
		"login−[object NavigatorLogin]",
		"userAgent−Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
		"userAgent−Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"mediaDevices−[object MediaDevices]",
		"locks−[object LockManager]",
		"webkitGetUserMedia−function webkitGetUserMedia() { [native code] }",
		"vendor−Google Inc.",
		"xr−[object XRSystem]",
		"mediaDevices−[object MediaDevices]",
		"virtualKeyboard−[object VirtualKeyboard]",
		"userAgent−Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"virtualKeyboard−[object VirtualKeyboard]",
		"appName−Netscape",
		"storageBuckets−[object StorageBucketManager]",
		"presentation−[object Presentation]",
		"onLine−true",
		"mimeTypes−[object MimeTypeArray]",
		"credentials−[object CredentialsContainer]",
		"presentation−[object Presentation]",
		"getGamepads−function getGamepads() { [native code] }",
		"vendorSub−",
		"virtualKeyboard−[object VirtualKeyboard]",
		"serviceWorker−[object ServiceWorkerContainer]",
		"xr−[object XRSystem]",
		"product−Gecko",
		"keyboard−[object Keyboard]",
		"gpu−[object GPU]",
		"getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
		"webkitPersistentStorage−[object DeprecatedStorageQuota]",
		"doNotTrack",
		"clearAppBadge−function clearAppBadge() { [native code] }",
		"presentation−[object Presentation]",
		"serial−[object Serial]",
		"locks−[object LockManager]",
		"requestMIDIAccess−function requestMIDIAccess() { [native code] }",
		"locks−[object LockManager]",
		"requestMediaKeySystemAccess−function requestMediaKeySystemAccess() { [native code] }",
		"vendor−Google Inc.",
		"pdfViewerEnabled−true",
		"language−zh-CN",
		"setAppBadge−function setAppBadge() { [native code] }",
		"geolocation−[object Geolocation]",
		"userAgentData−[object NavigatorUAData]",
		"mediaCapabilities−[object MediaCapabilities]",
		"requestMIDIAccess−function requestMIDIAccess() { [native code] }",
		"getUserMedia−function getUserMedia() { [native code] }",
		"mediaDevices−[object MediaDevices]",
		"webkitPersistentStorage−[object DeprecatedStorageQuota]",
		"userAgent−Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"sendBeacon−function sendBeacon() { [native code] }",
		"hardwareConcurrency−32",
		"appVersion−5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"credentials−[object CredentialsContainer]",
		"storage−[object StorageManager]",
		"cookieEnabled−true",
		"pdfViewerEnabled−true",
		"windowControlsOverlay−[object WindowControlsOverlay]",
		"scheduling−[object Scheduling]",
		"pdfViewerEnabled−true",
		"hardwareConcurrency−32",
		"xr−[object XRSystem]",
		"userAgent−Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		"webdriver−false",
		"getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
		"getInstalledRelatedApps−function getInstalledRelatedApps() { [native code] }",
		"bluetooth−[object Bluetooth]"}
	windowMap = []string{
		"0",
		"window",
		"self",
		"document",
		"name",
		"location",
		"customElements",
		"history",
		"navigation",
		"locationbar",
		"menubar",
		"personalbar",
		"scrollbars",
		"statusbar",
		"toolbar",
		"status",
		"closed",
		"frames",
		"length",
		"top",
		"opener",
		"parent",
		"frameElement",
		"navigator",
		"origin",
		"external",
		"screen",
		"innerWidth",
		"innerHeight",
		"scrollX",
		"pageXOffset",
		"scrollY",
		"pageYOffset",
		"visualViewport",
		"screenX",
		"screenY",
		"outerWidth",
		"outerHeight",
		"devicePixelRatio",
		"clientInformation",
		"screenLeft",
		"screenTop",
		"styleMedia",
		"onsearch",
		"isSecureContext",
		"trustedTypes",
		"performance",
		"onappinstalled",
		"onbeforeinstallprompt",
		"crypto",
		"indexedDB",
		"sessionStorage",
		"localStorage",
		"onbeforexrselect",
		"onabort",
		"onbeforeinput",
		"onbeforematch",
		"onbeforetoggle",
		"onblur",
		"oncancel",
		"oncanplay",
		"oncanplaythrough",
		"onchange",
		"onclick",
		"onclose",
		"oncontentvisibilityautostatechange",
		"oncontextlost",
		"oncontextmenu",
		"oncontextrestored",
		"oncuechange",
		"ondblclick",
		"ondrag",
		"ondragend",
		"ondragenter",
		"ondragleave",
		"ondragover",
		"ondragstart",
		"ondrop",
		"ondurationchange",
		"onemptied",
		"onended",
		"onerror",
		"onfocus",
		"onformdata",
		"oninput",
		"oninvalid",
		"onkeydown",
		"onkeypress",
		"onkeyup",
		"onload",
		"onloadeddata",
		"onloadedmetadata",
		"onloadstart",
		"onmousedown",
		"onmouseenter",
		"onmouseleave",
		"onmousemove",
		"onmouseout",
		"onmouseover",
		"onmouseup",
		"onmousewheel",
		"onpause",
		"onplay",
		"onplaying",
		"onprogress",
		"onratechange",
		"onreset",
		"onresize",
		"onscroll",
		"onsecuritypolicyviolation",
		"onseeked",
		"onseeking",
		"onselect",
		"onslotchange",
		"onstalled",
		"onsubmit",
		"onsuspend",
		"ontimeupdate",
		"ontoggle",
		"onvolumechange",
		"onwaiting",
		"onwebkitanimationend",
		"onwebkitanimationiteration",
		"onwebkitanimationstart",
		"onwebkittransitionend",
		"onwheel",
		"onauxclick",
		"ongotpointercapture",
		"onlostpointercapture",
		"onpointerdown",
		"onpointermove",
		"onpointerrawupdate",
		"onpointerup",
		"onpointercancel",
		"onpointerover",
		"onpointerout",
		"onpointerenter",
		"onpointerleave",
		"onselectstart",
		"onselectionchange",
		"onanimationend",
		"onanimationiteration",
		"onanimationstart",
		"ontransitionrun",
		"ontransitionstart",
		"ontransitionend",
		"ontransitioncancel",
		"onafterprint",
		"onbeforeprint",
		"onbeforeunload",
		"onhashchange",
		"onlanguagechange",
		"onmessage",
		"onmessageerror",
		"onoffline",
		"ononline",
		"onpagehide",
		"onpageshow",
		"onpopstate",
		"onrejectionhandled",
		"onstorage",
		"onunhandledrejection",
		"onunload",
		"crossOriginIsolated",
		"scheduler",
		"alert",
		"atob",
		"blur",
		"btoa",
		"cancelAnimationFrame",
		"cancelIdleCallback",
		"captureEvents",
		"clearInterval",
		"clearTimeout",
		"close",
		"confirm",
		"createImageBitmap",
		"fetch",
		"find",
		"focus",
		"getComputedStyle",
		"getSelection",
		"matchMedia",
		"moveBy",
		"moveTo",
		"open",
		"postMessage",
		"print",
		"prompt",
		"queueMicrotask",
		"releaseEvents",
		"reportError",
		"requestAnimationFrame",
		"requestIdleCallback",
		"resizeBy",
		"resizeTo",
		"scroll",
		"scrollBy",
		"scrollTo",
		"setInterval",
		"setTimeout",
		"stop",
		"structuredClone",
		"webkitCancelAnimationFrame",
		"webkitRequestAnimationFrame",
		"chrome",
		"g_opr",
		"opr",
		"ethereum",
		"caches",
		"cookieStore",
		"ondevicemotion",
		"ondeviceorientation",
		"ondeviceorientationabsolute",
		"launchQueue",
		"documentPictureInPicture",
		"getScreenDetails",
		"queryLocalFonts",
		"showDirectoryPicker",
		"showOpenFilePicker",
		"showSaveFilePicker",
		"originAgentCluster",
		"credentialless",
		"speechSynthesis",
		"onscrollend",
		"webkitRequestFileSystem",
		"webkitResolveLocalFileSystemURL",
		"__remixContext",
		"__oai_SSR_TTI",
		"__remixManifest",
		"__reactRouterVersion",
		"DD_RUM",
		"__REACT_INTL_CONTEXT__",
		"filterCSS",
		"filterXSS",
		"__SEGMENT_INSPECTOR__",
		"DD_LOGS",
		"regeneratorRuntime",
		"_g",
		"__remixRouteModules",
		"__remixRouter",
		"__STATSIG_SDK__",
		"__STATSIG_JS_SDK__",
		"__STATSIG_RERENDER_OVERRIDE__",
		"_oaiHandleSessionExpired"}
)
