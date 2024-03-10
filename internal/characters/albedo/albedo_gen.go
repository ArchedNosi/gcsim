// Code generated by "pipeline"; DO NOT EDIT.
package albedo

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
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
		attack_5,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.36739200353622437,
		0.3972960114479065,
		0.42719998955726624,
		0.46992000937461853,
		0.49982398748397827,
		0.5339999794960022,
		0.5809919834136963,
		0.6279839873313904,
		0.6749759912490845,
		0.7262399792671204,
		0.7849799990653992,
		0.854058027267456,
		0.9231359958648682,
		0.9922149777412415,
		1.0675729513168335,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.36739200353622437,
		0.3972960114479065,
		0.42719998955726624,
		0.46992000937461853,
		0.49982398748397827,
		0.5339999794960022,
		0.5809919834136963,
		0.6279839873313904,
		0.6749759912490845,
		0.7262399792671204,
		0.7849799990653992,
		0.854058027267456,
		0.9231359958648682,
		0.9922149777412415,
		1.0675729513168335,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.47454801201820374,
		0.5131739974021912,
		0.551800012588501,
		0.60698002576828,
		0.6456059813499451,
		0.6897500157356262,
		0.7504479885101318,
		0.8111460208892822,
		0.8718439936637878,
		0.9380599856376648,
		1.0139319896697998,
		1.103158950805664,
		1.192384958267212,
		1.2816109657287598,
		1.3789479732513428,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.49750998616218567,
		0.538004994392395,
		0.578499972820282,
		0.6363499760627747,
		0.6768450140953064,
		0.7231249809265137,
		0.7867599725723267,
		0.8503950238227844,
		0.9140300154685974,
		0.9834499955177307,
		1.0629940032958984,
		1.1565370559692383,
		1.250080943107605,
		1.3436239957809448,
		1.4456709623336792,
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.6207389831542969,
		0.6712650060653687,
		0.7217900156974792,
		0.7939689755439758,
		0.8444939851760864,
		0.9022380113601685,
		0.9816340208053589,
		1.0610309839248657,
		1.1404279470443726,
		1.2270430326461792,
		1.3262890577316284,
		1.4430030584335327,
		1.559715986251831,
		1.676429033279419,
		1.8037530183792114,
	}
	// attack: charge = [5 6]
	charge = [][]float64{
		{
			0.4729999899864197,
			0.5115000009536743,
			0.550000011920929,
			0.6050000190734863,
			0.6434999704360962,
			0.6875,
			0.7480000257492065,
			0.8084999918937683,
			0.8690000176429749,
			0.9350000023841858,
			1.0106250047683716,
			1.099560022354126,
			1.1884950399398804,
			1.2774300575256348,
			1.3744499683380127,
		},
		{
			0.6019999980926514,
			0.6510000228881836,
			0.699999988079071,
			0.7699999809265137,
			0.8190000057220459,
			0.875,
			0.9520000219345093,
			1.0290000438690186,
			1.1059999465942383,
			1.190000057220459,
			1.2862499952316284,
			1.399440050125122,
			1.5126299858093262,
			1.6258200407028198,
			1.7493000030517578,
		},
	}
	// attack: collision = [8]
	collision = []float64{
		0.6393240094184875,
		0.6913620233535767,
		0.743399977684021,
		0.8177400231361389,
		0.8697779774665833,
		0.9292500019073486,
		1.011023998260498,
		1.0927979946136475,
		1.1745719909667969,
		1.2637799978256226,
		1.3529880046844482,
		1.442196011543274,
		1.5314040184020996,
		1.6206120252609253,
		1.709820032119751,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		1.59676194190979,
		1.7267309427261353,
		1.8566999435424805,
		2.042370080947876,
		2.1723389625549316,
		2.3208749294281006,
		2.5251119136810303,
		2.72934889793396,
		2.9335858821868896,
		3.1563899517059326,
		3.3791940212249756,
		3.6019980907440186,
		3.8248019218444824,
		4.047605991363525,
		4.270410060882568,
	}
	// attack: lowPlunge = [9]
	lowPlunge = []float64{
		1.2783770561218262,
		1.3824310302734375,
		1.4864850044250488,
		1.635133981704712,
		1.7391870021820068,
		1.858106017112732,
		2.021620035171509,
		2.1851329803466797,
		2.3486459255218506,
		2.527024984359741,
		2.7054030895233154,
		2.8837809562683105,
		3.0621590614318848,
		3.24053692817688,
		3.418915033340454,
	}
	// skill: skill = [0]
	skill = []float64{
		1.3040000200271606,
		1.4018000364303589,
		1.4996000528335571,
		1.6299999952316284,
		1.7278000116348267,
		1.825600028038025,
		1.9559999704360962,
		2.086400032043457,
		2.2167999744415283,
		2.3471999168395996,
		2.47760009765625,
		2.6080000400543213,
		2.7709999084472656,
		2.934000015258789,
		3.0969998836517334,
	}
	// skill: skillTick = [1]
	skillTick = []float64{
		1.3359999656677246,
		1.4362000226974487,
		1.5363999605178833,
		1.6699999570846558,
		1.7702000141143799,
		1.8703999519348145,
		2.003999948501587,
		2.1375999450683594,
		2.271199941635132,
		2.4047999382019043,
		2.5383999347686768,
		2.671999931335449,
		2.8389999866485596,
		3.00600004196167,
		3.1730000972747803,
	}
	// burst: burst = [0]
	burst = []float64{
		3.671999931335449,
		3.9474000930786133,
		4.222799777984619,
		4.590000152587891,
		4.8653998374938965,
		5.1407999992370605,
		5.507999897003174,
		5.875199794769287,
		6.242400169372559,
		6.609600067138672,
		6.976799964904785,
		7.343999862670898,
		7.802999973297119,
		8.26200008392334,
		8.720999717712402,
	}
	// burst: burstPerBloom = [1]
	burstPerBloom = []float64{
		0.7200000286102295,
		0.7739999890327454,
		0.828000009059906,
		0.8999999761581421,
		0.9539999961853027,
		1.0080000162124634,
		1.0800000429153442,
		1.1519999504089355,
		1.2239999771118164,
		1.2960000038146973,
		1.3680000305175781,
		1.440000057220459,
		1.5299999713897705,
		1.6200000047683716,
		1.7100000381469727,
	}
)
