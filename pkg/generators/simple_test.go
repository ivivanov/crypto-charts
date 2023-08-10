package generators

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ivivanov/crypto-charts/pkg/fetchers/bitstamp"
)

const (
	OHLC_SAMPLE = `{"data": {"ohlc": [{"close": "30300", "high": "30300", "low": "30182", "open": "30203", "timestamp": "1689598800", "volume": "0.06869897"}, {"close": "30281", "high": "30294", "low": "30235", "open": "30284", "timestamp": "1689602400", "volume": "2.49486031"}, {"close": "30173", "high": "30292", "low": "30129", "open": "30291", "timestamp": "1689606000", "volume": "0.48458109"}, {"close": "30141", "high": "30234", "low": "30078", "open": "30181", "timestamp": "1689609600", "volume": "0.37175799"}, {"close": "30043", "high": "30152", "low": "29928", "open": "30139", "timestamp": "1689613200", "volume": "0.72399303"}, {"close": "29841", "high": "30042", "low": "29665", "open": "30042", "timestamp": "1689616800", "volume": "1.75529629"}, {"close": "29900", "high": "29925", "low": "29793", "open": "29840", "timestamp": "1689620400", "volume": "1.22594669"}, {"close": "29917", "high": "29956", "low": "29897", "open": "29897", "timestamp": "1689624000", "volume": "0.25695604"}, {"close": "30244", "high": "30280", "low": "29929", "open": "29929", "timestamp": "1689627600", "volume": "0.45502779"}, {"close": "30120", "high": "30263", "low": "30109", "open": "30263", "timestamp": "1689631200", "volume": "0.38141194"}, {"close": "30140", "high": "30158", "low": "30095", "open": "30095", "timestamp": "1689634800", "volume": "0.38453592"}, {"close": "30162", "high": "30198", "low": "30119", "open": "30141", "timestamp": "1689638400", "volume": "0.45551193"}, {"close": "30239", "high": "30239", "low": "30143", "open": "30172", "timestamp": "1689642000", "volume": "0.50083967"}, {"close": "30118", "high": "30238", "low": "30110", "open": "30222", "timestamp": "1689645600", "volume": "2.03308466"}, {"close": "30101", "high": "30151", "low": "30085", "open": "30114", "timestamp": "1689649200", "volume": "0.65263854"}, {"close": "30104", "high": "30108", "low": "30052", "open": "30079", "timestamp": "1689652800", "volume": "0.14687436"}, {"close": "30036", "high": "30116", "low": "30036", "open": "30114", "timestamp": "1689656400", "volume": "0.85234908"}, {"close": "29986", "high": "30092", "low": "29921", "open": "30060", "timestamp": "1689660000", "volume": "1.51926331"}, {"close": "29943", "high": "30030", "low": "29943", "open": "29978", "timestamp": "1689663600", "volume": "0.20467390"}, {"close": "30033", "high": "30077", "low": "29884", "open": "29965", "timestamp": "1689667200", "volume": "0.60684102"}, {"close": "29960", "high": "30043", "low": "29946", "open": "30043", "timestamp": "1689670800", "volume": "0.18502345"}, {"close": "30006", "high": "30026", "low": "29929", "open": "29977", "timestamp": "1689674400", "volume": "0.24419882"}, {"close": "29874", "high": "30049", "low": "29801", "open": "30003", "timestamp": "1689678000", "volume": "1.28155899"}, {"close": "29801", "high": "29910", "low": "29748", "open": "29878", "timestamp": "1689681600", "volume": "1.29531868"}, {"close": "29825", "high": "29856", "low": "29698", "open": "29782", "timestamp": "1689685200", "volume": "0.82725455"}, {"close": "29908", "high": "30014", "low": "29832", "open": "29836", "timestamp": "1689688800", "volume": "0.57302209"}, {"close": "29914", "high": "29929", "low": "29857", "open": "29901", "timestamp": "1689692400", "volume": "0.66772162"}, {"close": "29721", "high": "29957", "low": "29513", "open": "29898", "timestamp": "1689696000", "volume": "3.05011604"}, {"close": "29809", "high": "29809", "low": "29698", "open": "29698", "timestamp": "1689699600", "volume": "0.20078949"}, {"close": "29890", "high": "29919", "low": "29799", "open": "29826", "timestamp": "1689703200", "volume": "0.48734855"}, {"close": "29695", "high": "29906", "low": "29695", "open": "29906", "timestamp": "1689706800", "volume": "0.38643077"}, {"close": "29809", "high": "29809", "low": "29711", "open": "29732", "timestamp": "1689710400", "volume": "0.28250943"}, {"close": "29797", "high": "29818", "low": "29766", "open": "29773", "timestamp": "1689714000", "volume": "0.11623006"}, {"close": "29796", "high": "29826", "low": "29760", "open": "29788", "timestamp": "1689717600", "volume": "2.86092838"}, {"close": "29856", "high": "29864", "low": "29784", "open": "29813", "timestamp": "1689721200", "volume": "0.32721071"}, {"close": "30062", "high": "30062", "low": "29835", "open": "29854", "timestamp": "1689724800", "volume": "0.85386724"}, {"close": "29958", "high": "30042", "low": "29958", "open": "30038", "timestamp": "1689728400", "volume": "0.19545563"}, {"close": "30015", "high": "30021", "low": "29995", "open": "30018", "timestamp": "1689732000", "volume": "0.10566716"}, {"close": "30068", "high": "30071", "low": "30004", "open": "30032", "timestamp": "1689735600", "volume": "0.57102055"}, {"close": "30150", "high": "30192", "low": "30062", "open": "30074", "timestamp": "1689739200", "volume": "0.59529923"}, {"close": "30110", "high": "30110", "low": "30073", "open": "30095", "timestamp": "1689742800", "volume": "0.43598562"}, {"close": "30047", "high": "30087", "low": "30023", "open": "30070", "timestamp": "1689746400", "volume": "0.86141425"}, {"close": "30009", "high": "30063", "low": "29986", "open": "30047", "timestamp": "1689750000", "volume": "0.31220201"}, {"close": "29936", "high": "30005", "low": "29899", "open": "29983", "timestamp": "1689753600", "volume": "0.27016637"}, {"close": "30012", "high": "30036", "low": "29875", "open": "29886", "timestamp": "1689757200", "volume": "1.84008033"}, {"close": "29988", "high": "30034", "low": "29988", "open": "30002", "timestamp": "1689760800", "volume": "0.30546311"}, {"close": "30013", "high": "30029", "low": "29990", "open": "29997", "timestamp": "1689764400", "volume": "0.09483796"}, {"close": "29961", "high": "30020", "low": "29894", "open": "30017", "timestamp": "1689768000", "volume": "0.22813194"}, {"close": "30003", "high": "30133", "low": "29840", "open": "29961", "timestamp": "1689771600", "volume": "1.21836237"}, {"close": "29832", "high": "30056", "low": "29788", "open": "30056", "timestamp": "1689775200", "volume": "4.56677751"}, {"close": "29964", "high": "29986", "low": "29818", "open": "29818", "timestamp": "1689778800", "volume": "0.82445709"}, {"close": "29882", "high": "29913", "low": "29882", "open": "29909", "timestamp": "1689782400", "volume": "0.29251539"}, {"close": "29968", "high": "30090", "low": "29918", "open": "29918", "timestamp": "1689786000", "volume": "1.13941490"}, {"close": "30077", "high": "30077", "low": "29952", "open": "29992", "timestamp": "1689789600", "volume": "1.05143681"}, {"close": "30071", "high": "30095", "low": "30021", "open": "30055", "timestamp": "1689793200", "volume": "0.31188413"}, {"close": "29965", "high": "30044", "low": "29965", "open": "30044", "timestamp": "1689796800", "volume": "0.17502855"}, {"close": "29932", "high": "29965", "low": "29861", "open": "29964", "timestamp": "1689800400", "volume": "0.28536129"}, {"close": "29871", "high": "29919", "low": "29861", "open": "29892", "timestamp": "1689804000", "volume": "0.36256381"}, {"close": "29922", "high": "29946", "low": "29826", "open": "29869", "timestamp": "1689807600", "volume": "0.19369258"}, {"close": "29995", "high": "30005", "low": "29894", "open": "29909", "timestamp": "1689811200", "volume": "0.35696187"}, {"close": "30018", "high": "30018", "low": "29967", "open": "30012", "timestamp": "1689814800", "volume": "0.29190328"}, {"close": "29975", "high": "30044", "low": "29974", "open": "29998", "timestamp": "1689818400", "volume": "0.40079733"}, {"close": "29949", "high": "29986", "low": "29929", "open": "29982", "timestamp": "1689822000", "volume": "0.28075798"}, {"close": "29963", "high": "29970", "low": "29925", "open": "29956", "timestamp": "1689825600", "volume": "0.10297184"}, {"close": "30167", "high": "30167", "low": "29957", "open": "29957", "timestamp": "1689829200", "volume": "0.38748186"}, {"close": "30120", "high": "30181", "low": "30092", "open": "30146", "timestamp": "1689832800", "volume": "0.39112536"}, {"close": "30200", "high": "30236", "low": "30104", "open": "30136", "timestamp": "1689836400", "volume": "1.13954605"}, {"close": "30267", "high": "30300", "low": "30182", "open": "30218", "timestamp": "1689840000", "volume": "0.82410093"}, {"close": "30345", "high": "30388", "low": "30216", "open": "30273", "timestamp": "1689843600", "volume": "1.42892850"}, {"close": "30301", "high": "30420", "low": "30289", "open": "30335", "timestamp": "1689847200", "volume": "0.53260736"}, {"close": "30293", "high": "30293", "low": "30278", "open": "30287", "timestamp": "1689850800", "volume": "0.09579498"}, {"close": "30220", "high": "30310", "low": "30219", "open": "30286", "timestamp": "1689854400", "volume": "0.23183589"}, {"close": "30279", "high": "30279", "low": "30178", "open": "30239", "timestamp": "1689858000", "volume": "0.90282334"}, {"close": "29753", "high": "30238", "low": "29636", "open": "30224", "timestamp": "1689861600", "volume": "9.26597304"}, {"close": "29810", "high": "29835", "low": "29726", "open": "29743", "timestamp": "1689865200", "volume": "0.84321925"}, {"close": "29737", "high": "29874", "low": "29737", "open": "29807", "timestamp": "1689868800", "volume": "0.50770993"}, {"close": "29767", "high": "29812", "low": "29704", "open": "29735", "timestamp": "1689872400", "volume": "0.22830435"}, {"close": "29765", "high": "29770", "low": "29572", "open": "29749", "timestamp": "1689876000", "volume": "1.43514707"}, {"close": "29753", "high": "29753", "low": "29687", "open": "29687", "timestamp": "1689879600", "volume": "0.42439465"}, {"close": "29730", "high": "29770", "low": "29689", "open": "29770", "timestamp": "1689883200", "volume": "0.45294781"}, {"close": "29889", "high": "29889", "low": "29733", "open": "29733", "timestamp": "1689886800", "volume": "0.18225842"}, {"close": "29816", "high": "29865", "low": "29810", "open": "29865", "timestamp": "1689890400", "volume": "0.14199120"}, {"close": "29798", "high": "29823", "low": "29798", "open": "29809", "timestamp": "1689894000", "volume": "0.11781727"}, {"close": "29796", "high": "29834", "low": "29780", "open": "29806", "timestamp": "1689897600", "volume": "0.21299841"}, {"close": "29887", "high": "29889", "low": "29772", "open": "29772", "timestamp": "1689901200", "volume": "0.46270522"}, {"close": "29944", "high": "29944", "low": "29844", "open": "29860", "timestamp": "1689904800", "volume": "0.14835714"}, {"close": "29911", "high": "29937", "low": "29894", "open": "29915", "timestamp": "1689908400", "volume": "0.03885235"}, {"close": "29894", "high": "29894", "low": "29863", "open": "29863", "timestamp": "1689912000", "volume": "0.02956519"}, {"close": "29808", "high": "29837", "low": "29808", "open": "29832", "timestamp": "1689915600", "volume": "0.08166245"}, {"close": "29855", "high": "29857", "low": "29809", "open": "29822", "timestamp": "1689919200", "volume": "0.14941326"}, {"close": "29759", "high": "29853", "low": "29759", "open": "29853", "timestamp": "1689922800", "volume": "0.58940534"}, {"close": "29803", "high": "29814", "low": "29758", "open": "29758", "timestamp": "1689926400", "volume": "0.19710008"}, {"close": "29757", "high": "29838", "low": "29752", "open": "29784", "timestamp": "1689930000", "volume": "0.31446031"}, {"close": "29795", "high": "29818", "low": "29768", "open": "29773", "timestamp": "1689933600", "volume": "0.32918038"}, {"close": "29777", "high": "29816", "low": "29734", "open": "29792", "timestamp": "1689937200", "volume": "0.54180923"}, {"close": "29864", "high": "29898", "low": "29783", "open": "29783", "timestamp": "1689940800", "volume": "0.28679689"}, {"close": "29796", "high": "29906", "low": "29796", "open": "29860", "timestamp": "1689944400", "volume": "0.73027379"}, {"close": "29855", "high": "29875", "low": "29779", "open": "29805", "timestamp": "1689948000", "volume": "0.68624845"}, {"close": "29874", "high": "29900", "low": "29858", "open": "29900", "timestamp": "1689951600", "volume": "0.19816725"}, {"close": "29830", "high": "29868", "low": "29815", "open": "29865", "timestamp": "1689955200", "volume": "0.67958574"}, {"close": "29911", "high": "29931", "low": "29825", "open": "29825", "timestamp": "1689958800", "volume": "1.89502231"}, {"close": "29985", "high": "30030", "low": "29921", "open": "29921", "timestamp": "1689962400", "volume": "0.81153299"}, {"close": "29891", "high": "29917", "low": "29774", "open": "29917", "timestamp": "1689966000", "volume": "0.72240524"}, {"close": "29876", "high": "29892", "low": "29837", "open": "29845", "timestamp": "1689969600", "volume": "0.19574819"}, {"close": "29916", "high": "29948", "low": "29875", "open": "29883", "timestamp": "1689973200", "volume": "0.41258647"}, {"close": "29925", "high": "29957", "low": "29894", "open": "29894", "timestamp": "1689976800", "volume": "0.09562628"}, {"close": "29914", "high": "29931", "low": "29891", "open": "29931", "timestamp": "1689980400", "volume": "0.30372884"}, {"close": "29979", "high": "29979", "low": "29901", "open": "29907", "timestamp": "1689984000", "volume": "0.14616726"}, {"close": "29929", "high": "29987", "low": "29929", "open": "29987", "timestamp": "1689987600", "volume": "0.32406997"}, {"close": "29937", "high": "29964", "low": "29929", "open": "29939", "timestamp": "1689991200", "volume": "0.18200216"}, {"close": "29928", "high": "29945", "low": "29890", "open": "29912", "timestamp": "1689994800", "volume": "0.48454877"}, {"close": "29887", "high": "29929", "low": "29887", "open": "29924", "timestamp": "1689998400", "volume": "0.03362799"}, {"close": "29877", "high": "29897", "low": "29874", "open": "29882", "timestamp": "1690002000", "volume": "0.05588001"}, {"close": "29883", "high": "29883", "low": "29867", "open": "29867", "timestamp": "1690005600", "volume": "0.02513659"}, {"close": "29963", "high": "29969", "low": "29881", "open": "29881", "timestamp": "1690009200", "volume": "0.30351761"}, {"close": "29944", "high": "29976", "low": "29944", "open": "29949", "timestamp": "1690012800", "volume": "0.07156002"}, {"close": "29920", "high": "29943", "low": "29920", "open": "29929", "timestamp": "1690016400", "volume": "0.26196996"}, {"close": "29888", "high": "29914", "low": "29888", "open": "29914", "timestamp": "1690020000", "volume": "0.04558845"}, {"close": "29888", "high": "29901", "low": "29873", "open": "29901", "timestamp": "1690023600", "volume": "0.33418679"}, {"close": "29832", "high": "29863", "low": "29832", "open": "29857", "timestamp": "1690027200", "volume": "0.10099405"}, {"close": "29860", "high": "29862", "low": "29851", "open": "29852", "timestamp": "1690030800", "volume": "0.01817936"}, {"close": "29861", "high": "29889", "low": "29861", "open": "29877", "timestamp": "1690034400", "volume": "0.43279277"}, {"close": "29899", "high": "29899", "low": "29884", "open": "29884", "timestamp": "1690038000", "volume": "0.07467101"}, {"close": "29903", "high": "29903", "low": "29866", "open": "29879", "timestamp": "1690041600", "volume": "0.12642699"}, {"close": "29819", "high": "29890", "low": "29819", "open": "29888", "timestamp": "1690045200", "volume": "0.12402309"}, {"close": "29825", "high": "29828", "low": "29807", "open": "29818", "timestamp": "1690048800", "volume": "0.07720241"}, {"close": "29845", "high": "29848", "low": "29837", "open": "29837", "timestamp": "1690052400", "volume": "0.38751000"}, {"close": "29833", "high": "29834", "low": "29822", "open": "29832", "timestamp": "1690056000", "volume": "0.50211409"}, {"close": "29830", "high": "29834", "low": "29824", "open": "29833", "timestamp": "1690059600", "volume": "0.07978719"}, {"close": "29803", "high": "29826", "low": "29803", "open": "29826", "timestamp": "1690063200", "volume": "0.05769678"}, {"close": "29783", "high": "29798", "low": "29627", "open": "29797", "timestamp": "1690066800", "volume": "0.70932970"}, {"close": "29771", "high": "29797", "low": "29771", "open": "29784", "timestamp": "1690070400", "volume": "0.32275003"}, {"close": "29850", "high": "29874", "low": "29794", "open": "29794", "timestamp": "1690074000", "volume": "0.82605587"}, {"close": "29889", "high": "29889", "low": "29817", "open": "29817", "timestamp": "1690077600", "volume": "1.40004539"}, {"close": "29853", "high": "29866", "low": "29849", "open": "29866", "timestamp": "1690081200", "volume": "0.05287786"}, {"close": "29907", "high": "29911", "low": "29847", "open": "29847", "timestamp": "1690084800", "volume": "0.35681471"}, {"close": "29915", "high": "29960", "low": "29908", "open": "29908", "timestamp": "1690088400", "volume": "0.22435282"}, {"close": "29919", "high": "29934", "low": "29900", "open": "29934", "timestamp": "1690092000", "volume": "0.17876409"}, {"close": "29886", "high": "29892", "low": "29886", "open": "29892", "timestamp": "1690095600", "volume": "0.00399727"}, {"close": "29950", "high": "29950", "low": "29934", "open": "29937", "timestamp": "1690099200", "volume": "0.05034914"}, {"close": "29951", "high": "29951", "low": "29943", "open": "29947", "timestamp": "1690102800", "volume": "0.01433011"}, {"close": "29903", "high": "29918", "low": "29903", "open": "29913", "timestamp": "1690106400", "volume": "0.02559675"}, {"close": "29902", "high": "29902", "low": "29895", "open": "29895", "timestamp": "1690110000", "volume": "0.00402413"}, {"close": "29900", "high": "29916", "low": "29900", "open": "29914", "timestamp": "1690113600", "volume": "0.03037667"}, {"close": "29888", "high": "29909", "low": "29860", "open": "29909", "timestamp": "1690117200", "volume": "1.21233904"}, {"close": "29902", "high": "29902", "low": "29888", "open": "29888", "timestamp": "1690120800", "volume": "0.99824640"}, {"close": "29902", "high": "29912", "low": "29900", "open": "29901", "timestamp": "1690124400", "volume": "0.24359822"}, {"close": "29952", "high": "29962", "low": "29916", "open": "29921", "timestamp": "1690128000", "volume": "0.06768829"}, {"close": "30100", "high": "30100", "low": "29966", "open": "29975", "timestamp": "1690131600", "volume": "0.28320917"}, {"close": "30273", "high": "30273", "low": "30095", "open": "30126", "timestamp": "1690135200", "volume": "1.70324236"}, {"close": "30091", "high": "30340", "low": "30063", "open": "30254", "timestamp": "1690138800", "volume": "0.50806091"}, {"close": "30107", "high": "30142", "low": "30107", "open": "30119", "timestamp": "1690142400", "volume": "0.08948605"}, {"close": "29970", "high": "30079", "low": "29938", "open": "30071", "timestamp": "1690146000", "volume": "0.34833468"}, {"close": "30019", "high": "30027", "low": "29966", "open": "29966", "timestamp": "1690149600", "volume": "0.23141812"}, {"close": "30085", "high": "30085", "low": "30037", "open": "30058", "timestamp": "1690153200", "volume": "0.25640650"}, {"close": "30022", "high": "30101", "low": "30022", "open": "30101", "timestamp": "1690156800", "volume": "0.27421728"}, {"close": "29909", "high": "29972", "low": "29882", "open": "29960", "timestamp": "1690160400", "volume": "2.12847585"}, {"close": "29788", "high": "29876", "low": "29767", "open": "29876", "timestamp": "1690164000", "volume": "0.42250525"}, {"close": "29693", "high": "29750", "low": "29693", "open": "29750", "timestamp": "1690167600", "volume": "0.22014909"}, {"close": "29775", "high": "29795", "low": "29667", "open": "29702", "timestamp": "1690171200", "volume": "0.06628435"}, {"close": "29796", "high": "29796", "low": "29770", "open": "29777", "timestamp": "1690174800", "volume": "2.02714293"}, {"close": "29794", "high": "29794", "low": "29775", "open": "29792", "timestamp": "1690178400", "volume": "0.15502677"}, {"close": "29821", "high": "29821", "low": "29735", "open": "29765", "timestamp": "1690182000", "volume": "0.01594350"}, {"close": "29757", "high": "29819", "low": "29757", "open": "29819", "timestamp": "1690185600", "volume": "0.37631531"}, {"close": "29221", "high": "29756", "low": "29001", "open": "29755", "timestamp": "1690189200", "volume": "7.72786100"}, {"close": "29324", "high": "29347", "low": "29140", "open": "29171", "timestamp": "1690192800", "volume": "1.94602418"}, {"close": "29250", "high": "29331", "low": "29250", "open": "29331", "timestamp": "1690196400", "volume": "0.59117491"}, {"close": "29223", "high": "29224", "low": "29164", "open": "29219", "timestamp": "1690200000", "volume": "3.14410635"}], "pair": "BTC/USDT"}}`
)

func TestNewLineChart(t *testing.T) {
	lineChartGenerator := &SimpleLineChartGenerator{}

	ohlcData := map[string]bitstamp.OHLCData{}
	err := json.Unmarshal([]byte(OHLC_SAMPLE), &ohlcData)
	if err != nil {
		t.Fatal(err)
	}

	ohlc := ohlcData["data"].OHCL
	historicalData := bitstamp.MapOHLCtoMarketInfo(ohlc)
	svg, err := lineChartGenerator.NewLineChart(historicalData)
	if err != nil {
		t.Fatal(err)
	}

	expIndex := 0
	expString := "<svg"
	actIndex := strings.Index(svg, expString)
	if actIndex != expIndex {
		t.Fatalf("exp: %v, act: %v", expIndex, actIndex)
	}

	errorStr := "cannot draw values"
	if strings.Contains(svg, errorStr) {
		t.Fatal("error generating svg")
	}
}
