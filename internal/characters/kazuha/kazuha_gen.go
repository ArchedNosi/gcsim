// Code generated by "pipeline"; DO NOT EDIT.
package kazuha

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		attack_3,
		{attack_4},
		attack_5,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.4497799873352051,
		0.48638999462127686,
		0.5230000019073486,
		0.5752999782562256,
		0.6119099855422974,
		0.6537500023841858,
		0.7112799882888794,
		0.768809974193573,
		0.8263400197029114,
		0.8891000151634216,
		0.9610130190849304,
		1.0455820560455322,
		1.1301510334014893,
		1.2147200107574463,
		1.3069770336151123,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.45236000418663025,
		0.4891799986362457,
		0.5260000228881836,
		0.5785999894142151,
		0.6154199838638306,
		0.6575000286102295,
		0.7153599858283997,
		0.7732200026512146,
		0.8310800194740295,
		0.8942000269889832,
		0.9665250182151794,
		1.0515789985656738,
		1.136633038520813,
		1.2216880321502686,
		1.3144739866256714,
	}
	// attack: attack_3 = [2 3]
	attack_3 = [][]float64{
		{
			0.257999986410141,
			0.27900001406669617,
			0.30000001192092896,
			0.33000001311302185,
			0.35100001096725464,
			0.375,
			0.40799999237060547,
			0.4410000145435333,
			0.4740000069141388,
			0.5099999904632568,
			0.5512499809265137,
			0.5997599959373474,
			0.6482700109481812,
			0.6967800259590149,
			0.7497000098228455,
		},
		{
			0.30959999561309814,
			0.33480000495910645,
			0.36000001430511475,
			0.3959999978542328,
			0.4212000072002411,
			0.44999998807907104,
			0.4896000027656555,
			0.52920001745224,
			0.5687999725341797,
			0.6119999885559082,
			0.6614999771118164,
			0.7197120189666748,
			0.7779240012168884,
			0.836135983467102,
			0.8996400237083435,
		},
	}
	// attack: attack_4 = [4]
	attack_4 = []float64{
		0.6071599721908569,
		0.6565799713134766,
		0.7059999704360962,
		0.7766000032424927,
		0.8260200023651123,
		0.8824999928474426,
		0.9601600170135498,
		1.0378199815750122,
		1.1154799461364746,
		1.2001999616622925,
		1.2972749471664429,
		1.4114350080490112,
		1.52559494972229,
		1.6397559642791748,
		1.7642940282821655,
	}
	// attack: attack_5 = [5 5 5]
	attack_5 = [][]float64{
		{
			0.25369998812675476,
			0.27434998750686646,
			0.29499998688697815,
			0.3244999945163727,
			0.3451499938964844,
			0.3687500059604645,
			0.40119999647140503,
			0.4336499869823456,
			0.4661000072956085,
			0.5015000104904175,
			0.5420629978179932,
			0.5897639989852905,
			0.6374650001525879,
			0.6851670145988464,
			0.7372050285339355,
		},
		{
			0.25369998812675476,
			0.27434998750686646,
			0.29499998688697815,
			0.3244999945163727,
			0.3451499938964844,
			0.3687500059604645,
			0.40119999647140503,
			0.4336499869823456,
			0.4661000072956085,
			0.5015000104904175,
			0.5420629978179932,
			0.5897639989852905,
			0.6374650001525879,
			0.6851670145988464,
			0.7372050285339355,
		},
		{
			0.25369998812675476,
			0.27434998750686646,
			0.29499998688697815,
			0.3244999945163727,
			0.3451499938964844,
			0.3687500059604645,
			0.40119999647140503,
			0.4336499869823456,
			0.4661000072956085,
			0.5015000104904175,
			0.5420629978179932,
			0.5897639989852905,
			0.6374650001525879,
			0.6851670145988464,
			0.7372050285339355,
		},
	}
	// attack: charge = [6 7]
	charge = [][]float64{
		{
			0.4300000071525574,
			0.4650000035762787,
			0.5,
			0.550000011920929,
			0.5849999785423279,
			0.625,
			0.6800000071525574,
			0.7350000143051147,
			0.7900000214576721,
			0.8500000238418579,
			0.918749988079071,
			0.9995999932289124,
			1.0804500579833984,
			1.1612999439239502,
			1.249500036239624,
		},
		{
			0.7464799880981445,
			0.8072400093078613,
			0.8679999709129333,
			0.954800009727478,
			1.0155600309371948,
			1.0850000381469727,
			1.1804800033569336,
			1.2759599685668945,
			1.371440052986145,
			1.475600004196167,
			1.5949499607086182,
			1.7353060245513916,
			1.875661015510559,
			2.016016960144043,
			2.1691319942474365,
		},
	}
	// attack: collision = [9]
	collision = []float64{
		0.8183349967002869,
		0.8849430084228516,
		0.9515519738197327,
		1.046707034111023,
		1.1133160591125488,
		1.1894400119781494,
		1.2941110134124756,
		1.3987809419631958,
		1.503451943397522,
		1.6176379919052124,
		1.7318249940872192,
		1.8460110425949097,
		1.9601969718933105,
		2.074383020401001,
		2.188570022583008,
	}
	// attack: highPlunge = [11]
	highPlunge = []float64{
		2.0438549518585205,
		2.2102160453796387,
		2.3765759468078613,
		2.61423397064209,
		2.7805941104888916,
		2.970720052719116,
		3.232142925262451,
		3.4935669898986816,
		3.7549901008605957,
		4.0401787757873535,
		4.3253679275512695,
		4.6105570793151855,
		4.895747184753418,
		5.180935859680176,
		5.466125011444092,
	}
	// attack: lowPlunge = [10]
	lowPlunge = []float64{
		1.6363229751586914,
		1.7695120573043823,
		1.9027010202407837,
		2.092971086502075,
		2.2261600494384766,
		2.378376007080078,
		2.5876729488372803,
		2.7969698905944824,
		3.0062670707702637,
		3.234591007232666,
		3.4629149436950684,
		3.691240072250366,
		3.9195640087127686,
		4.14788818359375,
		4.376212120056152,
	}
	// skill: skill = [0]
	skill = []float64{
		1.9199999570846558,
		2.063999891281128,
		2.2079999446868896,
		2.4000000953674316,
		2.5439999103546143,
		2.687999963760376,
		2.880000114440918,
		3.072000026702881,
		3.2639999389648438,
		3.4560000896453857,
		3.6480000019073486,
		3.8399999141693115,
		4.079999923706055,
		4.320000171661377,
		4.559999942779541,
	}
	// skill: skillHold = [2]
	skillHold = []float64{
		2.6080000400543213,
		2.8036000728607178,
		2.9992001056671143,
		3.259999990463257,
		3.4556000232696533,
		3.65120005607605,
		3.9119999408721924,
		4.172800064086914,
		4.433599948883057,
		4.694399833679199,
		4.9552001953125,
		5.216000080108643,
		5.541999816894531,
		5.868000030517578,
		6.193999767303467,
	}
	// burst: burstDot = [1]
	burstDot = []float64{
		1.2000000476837158,
		1.2899999618530273,
		1.3799999952316284,
		1.5,
		1.590000033378601,
		1.6799999475479126,
		1.7999999523162842,
		1.9199999570846558,
		2.0399999618530273,
		2.1600000858306885,
		2.2799999713897705,
		2.4000000953674316,
		2.549999952316284,
		2.700000047683716,
		2.8499999046325684,
	}
	// burst: burstEleDot = [2]
	burstEleDot = []float64{
		0.36000001430511475,
		0.3869999945163727,
		0.414000004529953,
		0.44999998807907104,
		0.47699999809265137,
		0.5040000081062317,
		0.5400000214576721,
		0.5759999752044678,
		0.6119999885559082,
		0.6480000019073486,
		0.6840000152587891,
		0.7200000286102295,
		0.7649999856948853,
		0.8100000023841858,
		0.8550000190734863,
	}
	// burst: burstSlash = [0]
	burstSlash = []float64{
		2.624000072479248,
		2.8208000659942627,
		3.0176000595092773,
		3.2799999713897705,
		3.476799964904785,
		3.6735999584198,
		3.936000108718872,
		4.198400020599365,
		4.4608001708984375,
		4.723199844360352,
		4.985599994659424,
		5.248000144958496,
		5.576000213623047,
		5.9039998054504395,
		6.23199987411499,
	}
)
