package mt_test

import (
	"fmt"
	"github.com/nasa9084/go-mersenne-twister"
	"testing"
)

var expectedInt = []uint32{
	1067595299, 955945823, 477289528, 4107218783, 4228976476,
	3344332714, 3355579695, 227628506, 810200273, 2591290167,
	2560260675, 3242736208, 646746669, 1479517882, 4245472273,
	1143372638, 3863670494, 3221021970, 1773610557, 1138697238,
	1421897700, 1269916527, 2859934041, 1764463362, 3874892047,
	3965319921, 72549643, 2383988930, 2600218693, 3237492380,
	2792901476, 725331109, 605841842, 271258942, 715137098,
	3297999536, 1322965544, 4229579109, 1395091102, 3735697720,
	2101727825, 3730287744, 2950434330, 1661921839, 2895579582,
	2370511479, 1004092106, 2247096681, 2111242379, 3237345263,
	4082424759, 219785033, 2454039889, 3709582971, 835606218,
	2411949883, 2735205030, 756421180, 2175209704, 1873865952,
	2762534237, 4161807854, 3351099340, 181129879, 3269891896,
	776029799, 2218161979, 3001745796, 1866825872, 2133627728,
	34862734, 1191934573, 3102311354, 2916517763, 1012402762,
	2184831317, 4257399449, 2899497138, 3818095062, 3030756734,
	1282161629, 420003642, 2326421477, 2741455717, 1278020671,
	3744179621, 271777016, 2626330018, 2560563991, 3055977700,
	4233527566, 1228397661, 3595579322, 1077915006, 2395931898,
	1851927286, 3013683506, 1999971931, 3006888962, 1049781534,
	1488758959, 3491776230, 104418065, 2448267297, 3075614115,
	3872332600, 891912190, 3936547759, 2269180963, 2633455084,
	1047636807, 2604612377, 2709305729, 1952216715, 207593580,
	2849898034, 670771757, 2210471108, 467711165, 263046873,
	3569667915, 1042291111, 3863517079, 1464270005, 2758321352,
	3790799816, 2301278724, 3106281430, 7974801, 2792461636,
	555991332, 621766759, 1322453093, 853629228, 686962251,
	1455120532, 957753161, 1802033300, 1021534190, 3486047311,
	1902128914, 3701138056, 4176424663, 1795608698, 560858864,
	3737752754, 3141170998, 1553553385, 3367807274, 711546358,
	2475125503, 262969859, 251416325, 2980076994, 1806565895,
	969527843, 3529327173, 2736343040, 2987196734, 1649016367,
	2206175811, 3048174801, 3662503553, 3138851612, 2660143804,
	1663017612, 1816683231, 411916003, 3887461314, 2347044079,
	1015311755, 1203592432, 2170947766, 2569420716, 813872093,
	1105387678, 1431142475, 220570551, 4243632715, 4179591855,
	2607469131, 3090613241, 282341803, 1734241730, 1391822177,
	1001254810, 827927915, 1886687171, 3935097347, 2631788714,
	3905163266, 110554195, 2447955646, 3717202975, 3304793075,
	3739614479, 3059127468, 953919171, 2590123714, 1132511021,
	3795593679, 2788030429, 982155079, 3472349556, 859942552,
	2681007391, 2299624053, 647443547, 233600422, 608168955,
	3689327453, 1849778220, 1608438222, 3968158357, 2692977776,
	2851872572, 246750393, 3582818628, 3329652309, 4036366910,
	1012970930, 950780808, 3959768744, 2538550045, 191422718,
	2658142375, 3276369011, 2927737484, 1234200027, 1920815603,
	3536074689, 1535612501, 2184142071, 3276955054, 428488088,
	2378411984, 4059769550, 3913744741, 2732139246, 64369859,
	3755670074, 842839565, 2819894466, 2414718973, 1010060670,
	1839715346, 2410311136, 152774329, 3485009480, 4102101512,
	2852724304, 879944024, 1785007662, 2748284463, 1354768064,
	3267784736, 2269127717, 3001240761, 3179796763, 895723219,
	865924942, 4291570937, 89355264, 1471026971, 4114180745,
	3201939751, 2867476999, 2460866060, 3603874571, 2238880432,
	3308416168, 2072246611, 2755653839, 3773737248, 1709066580,
	4282731467, 2746170170, 2832568330, 433439009, 3175778732,
	26248366, 2551382801, 183214346, 3893339516, 1928168445,
	1337157619, 3429096554, 3275170900, 1782047316, 4264403756,
	1876594403, 4289659572, 3223834894, 1728705513, 4068244734,
	2867840287, 1147798696, 302879820, 1730407747, 1923824407,
	1180597908, 1569786639, 198796327, 560793173, 2107345620,
	2705990316, 3448772106, 3678374155, 758635715, 884524671,
	486356516, 1774865603, 3881226226, 2635213607, 1181121587,
	1508809820, 3178988241, 1594193633, 1235154121, 326117244,
	2304031425, 937054774, 2687415945, 3192389340, 2003740439,
	1823766188, 2759543402, 10067710, 1533252662, 4132494984,
	82378136, 420615890, 3467563163, 541562091, 3535949864,
	2277319197, 3330822853, 3215654174, 4113831979, 4204996991,
	2162248333, 3255093522, 2219088909, 2978279037, 255818579,
	2859348628, 3097280311, 2569721123, 1861951120, 2907080079,
	2719467166, 998319094, 2521935127, 2404125338, 259456032,
	2086860995, 1839848496, 1893547357, 2527997525, 1489393124,
	2860855349, 76448234, 2264934035, 744914583, 2586791259,
	1385380501, 66529922, 1819103258, 1899300332, 2098173828,
	1793831094, 276463159, 360132945, 4178212058, 595015228,
	177071838, 2800080290, 1573557746, 1548998935, 378454223,
	1460534296, 1116274283, 3112385063, 3709761796, 827999348,
	3580042847, 1913901014, 614021289, 4278528023, 1905177404,
	45407939, 3298183234, 1184848810, 3644926330, 3923635459,
	1627046213, 3677876759, 969772772, 1160524753, 1522441192,
	452369933, 1527502551, 832490847, 1003299676, 1071381111,
	2891255476, 973747308, 4086897108, 1847554542, 3895651598,
	2227820339, 1621250941, 2881344691, 3583565821, 3510404498,
	849362119, 862871471, 797858058, 2867774932, 2821282612,
	3272403146, 3997979905, 209178708, 1805135652, 6783381,
	2823361423, 792580494, 4263749770, 776439581, 3798193823,
	2853444094, 2729507474, 1071873341, 1329010206, 1289336450,
	3327680758, 2011491779, 80157208, 922428856, 1158943220,
	1667230961, 2461022820, 2608845159, 387516115, 3345351910,
	1495629111, 4098154157, 3156649613, 3525698599, 4134908037,
	446713264, 2137537399, 3617403512, 813966752, 1157943946,
	3734692965, 1680301658, 3180398473, 3509854711, 2228114612,
	1008102291, 486805123, 863791847, 3189125290, 1050308116,
	3777341526, 4291726501, 844061465, 1347461791, 2826481581,
	745465012, 2055805750, 4260209475, 2386693097, 2980646741,
	447229436, 2077782664, 1232942813, 4023002732, 1399011509,
	3140569849, 2579909222, 3794857471, 900758066, 2887199683,
	1720257997, 3367494931, 2668921229, 955539029, 3818726432,
	1105704962, 3889207255, 2277369307, 2746484505, 1761846513,
	2413916784, 2685127085, 4240257943, 1166726899, 4215215715,
	3082092067, 3960461946, 1663304043, 2087473241, 4162589986,
	2507310778, 1579665506, 767234210, 970676017, 492207530,
	1441679602, 1314785090, 3262202570, 3417091742, 1561989210,
	3011406780, 1146609202, 3262321040, 1374872171, 1634688712,
	1280458888, 2230023982, 419323804, 3262899800, 39783310,
	1641619040, 1700368658, 2207946628, 2571300939, 2424079766,
	780290914, 2715195096, 3390957695, 163151474, 2309534542,
	1860018424, 555755123, 280320104, 1604831083, 2713022383,
	1728987441, 3639955502, 623065489, 3828630947, 4275479050,
	3516347383, 2343951195, 2430677756, 635534992, 3868699749,
	808442435, 3070644069, 4282166003, 2093181383, 2023555632,
	1568662086, 3422372620, 4134522350, 3016979543, 3259320234,
	2888030729, 3185253876, 4258779643, 1267304371, 1022517473,
	815943045, 929020012, 2995251018, 3371283296, 3608029049,
	2018485115, 122123397, 2810669150, 1411365618, 1238391329,
	1186786476, 3155969091, 2242941310, 1765554882, 279121160,
	4279838515, 1641578514, 3796324015, 13351065, 103516986,
	1609694427, 551411743, 2493771609, 1316337047, 3932650856,
	4189700203, 463397996, 2937735066, 1855616529, 2626847990,
	55091862, 3823351211, 753448970, 4045045500, 1274127772,
	1124182256, 92039808, 2126345552, 425973257, 386287896,
	2589870191, 1987762798, 4084826973, 2172456685, 3366583455,
	3602966653, 2378803535, 2901764433, 3716929006, 3710159000,
	2653449155, 3469742630, 3096444476, 3932564653, 2595257433,
	318974657, 3146202484, 853571438, 144400272, 3768408841,
	782634401, 2161109003, 570039522, 1886241521, 14249488,
	2230804228, 1604941699, 3928713335, 3921942509, 2155806892,
	134366254, 430507376, 1924011722, 276713377, 196481886,
	3614810992, 1610021185, 1785757066, 851346168, 3761148643,
	2918835642, 3364422385, 3012284466, 3735958851, 2643153892,
	3778608231, 1164289832, 205853021, 2876112231, 3503398282,
	3078397001, 3472037921, 1748894853, 2740861475, 316056182,
	1660426908, 168885906, 956005527, 3984354789, 566521563,
	1001109523, 1216710575, 2952284757, 3834433081, 3842608301,
	2467352408, 3974441264, 3256601745, 1409353924, 1329904859,
	2307560293, 3125217879, 3622920184, 3832785684, 3882365951,
	2308537115, 2659155028, 1450441945, 3532257603, 3186324194,
	1225603425, 1124246549, 175808705, 3009142319, 2796710159,
	3651990107, 160762750, 1902254979, 1698648476, 1134980669,
	497144426, 3302689335, 4057485630, 3603530763, 4087252587,
	427812652, 286876201, 823134128, 1627554964, 3745564327,
	2589226092, 4202024494, 62878473, 3275585894, 3987124064,
	2791777159, 1916869511, 2585861905, 1375038919, 1403421920,
	60249114, 3811870450, 3021498009, 2612993202, 528933105,
	2757361321, 3341402964, 2621861700, 273128190, 4015252178,
	3094781002, 1621621288, 2337611177, 1796718448, 1258965619,
	4241913140, 2138560392, 3022190223, 4174180924, 450094611,
	3274724580, 617150026, 2704660665, 1469700689, 1341616587,
	356715071, 1188789960, 2278869135, 1766569160, 2795896635,
	57824704, 2893496380, 1235723989, 1630694347, 3927960522,
	428891364, 1814070806, 2287999787, 4125941184, 3968103889,
	3548724050, 1025597707, 1404281500, 2002212197, 92429143,
	2313943944, 2403086080, 3006180634, 3561981764, 1671860914,
	1768520622, 1803542985, 844848113, 3006139921, 1410888995,
	1157749833, 2125704913, 1789979528, 1799263423, 741157179,
	2405862309, 767040434, 2655241390, 3663420179, 2172009096,
	2511931187, 1680542666, 231857466, 1154981000, 157168255,
	1454112128, 3505872099, 1929775046, 2309422350, 2143329496,
	2960716902, 407610648, 2938108129, 2581749599, 538837155,
	2342628867, 430543915, 740188568, 1937713272, 3315215132,
	2085587024, 4030765687, 766054429, 3517641839, 689721775,
	1294158986, 1753287754, 4202601348, 1974852792, 33459103,
	3568087535, 3144677435, 1686130825, 4134943013, 3005738435,
	3599293386, 426570142, 754104406, 3660892564, 1964545167,
	829466833, 821587464, 1746693036, 1006492428, 1595312919,
	1256599985, 1024482560, 1897312280, 2902903201, 691790057,
	1037515867, 3176831208, 1968401055, 2173506824, 1089055278,
	1748401123, 2941380082, 968412354, 1818753861, 2973200866,
	3875951774, 1119354008, 3988604139, 1647155589, 2232450826,
	3486058011, 3655784043, 3759258462, 847163678, 1082052057,
	989516446, 2871541755, 3196311070, 3929963078, 658187585,
	3664944641, 2175149170, 2203709147, 2756014689, 2456473919,
	3890267390, 1293787864, 2830347984, 3059280931, 4158802520,
	1561677400, 2586570938, 783570352, 1355506163, 31495586,
	3789437343, 3340549429, 2092501630, 896419368, 671715824,
	3530450081, 3603554138, 1055991716, 3442308219, 1499434728,
	3130288473, 3639507000, 17769680, 2259741420, 487032199,
	4227143402, 3693771256, 1880482820, 3924810796, 381462353,
	4017855991, 2452034943, 2736680833, 2209866385, 2128986379,
	437874044, 595759426, 641721026, 1636065708, 3899136933,
	629879088, 3591174506, 351984326, 2638783544, 2348444281,
	2341604660, 2123933692, 143443325, 1525942256, 364660499,
	599149312, 939093251, 1523003209, 106601097, 376589484,
	1346282236, 1297387043, 764598052, 3741218111, 933457002,
	1886424424, 3219631016, 525405256, 3014235619, 323149677,
	2038881721, 4100129043, 2851715101, 2984028078, 1888574695,
	2014194741, 3515193880, 4180573530, 3461824363, 2641995497,
	3179230245, 2902294983, 2217320456, 4040852155, 1784656905,
	3311906931, 87498458, 2752971818, 2635474297, 2831215366,
	3682231106, 2920043893, 3772929704, 2816374944, 309949752,
	2383758854, 154870719, 385111597, 1191604312, 1840700563,
	872191186, 2925548701, 1310412747, 2102066999, 1504727249,
	3574298750, 1191230036, 3330575266, 3180292097, 3539347721,
	681369118, 3305125752, 3648233597, 950049240, 4173257693,
	1760124957, 512151405, 681175196, 580563018, 1169662867,
	4015033554, 2687781101, 699691603, 2673494188, 1137221356,
	123599888, 472658308, 1053598179, 1012713758, 3481064843,
	3759461013, 3981457956, 3830587662, 1877191791, 3650996736,
	988064871, 3515461600, 4089077232, 2225147448, 1249609188,
	2643151863, 3896204135, 2416995901, 1397735321, 3460025646,
}

var expectedFloat = []float64{
	0.76275443, 0.99000644, 0.98670464, 0.10143112, 0.27933125,
	0.69867227, 0.94218740, 0.03427201, 0.78842173, 0.28180608,
	0.92179002, 0.20785655, 0.54534773, 0.69644020, 0.38107718,
	0.23978165, 0.65286910, 0.07514568, 0.22765211, 0.94872929,
	0.74557914, 0.62664415, 0.54708246, 0.90959343, 0.42043116,
	0.86334511, 0.19189126, 0.14718544, 0.70259889, 0.63426346,
	0.77408121, 0.04531601, 0.04605807, 0.88595519, 0.69398270,
	0.05377184, 0.61711170, 0.05565708, 0.10133577, 0.41500776,
	0.91810699, 0.22320679, 0.23353705, 0.92871862, 0.98897234,
	0.19786706, 0.80558809, 0.06961067, 0.55840445, 0.90479405,
	0.63288060, 0.95009721, 0.54948447, 0.20645042, 0.45000959,
	0.87050869, 0.70806991, 0.19406895, 0.79286390, 0.49332866,
	0.78483914, 0.75145146, 0.12341941, 0.42030252, 0.16728160,
	0.59906494, 0.37575460, 0.97815160, 0.39815952, 0.43595080,
	0.04952478, 0.33917805, 0.76509902, 0.61034321, 0.90654701,
	0.92915732, 0.85365931, 0.18812377, 0.65913428, 0.28814566,
	0.59476081, 0.27835931, 0.60722542, 0.68310435, 0.69387186,
	0.03699800, 0.65897714, 0.17527003, 0.02889304, 0.86777366,
	0.12352068, 0.91439461, 0.32022990, 0.44445731, 0.34903686,
	0.74639273, 0.65918367, 0.92492794, 0.31872642, 0.77749724,
	0.85413832, 0.76385624, 0.32744211, 0.91326300, 0.27458185,
	0.22190155, 0.19865383, 0.31227402, 0.85321225, 0.84243342,
	0.78544200, 0.71854080, 0.92503892, 0.82703064, 0.88306297,
	0.47284073, 0.70059042, 0.48003761, 0.38671694, 0.60465770,
	0.41747204, 0.47163243, 0.72750808, 0.65830223, 0.10955369,
	0.64215401, 0.23456345, 0.95944940, 0.72822249, 0.40888451,
	0.69980355, 0.26677428, 0.57333635, 0.39791582, 0.85377858,
	0.76962816, 0.72004885, 0.90903087, 0.51376506, 0.37732665,
	0.12691640, 0.71249738, 0.81217908, 0.37037313, 0.32772374,
	0.14238259, 0.05614811, 0.74363008, 0.39773267, 0.94859135,
	0.31452454, 0.11730313, 0.62962618, 0.33334237, 0.45547255,
	0.10089665, 0.56550662, 0.60539371, 0.16027624, 0.13245301,
	0.60959939, 0.04671662, 0.99356286, 0.57660859, 0.40269560,
	0.45274629, 0.06699735, 0.85064246, 0.87742744, 0.54508392,
	0.87242982, 0.29321385, 0.67660627, 0.68230715, 0.79052073,
	0.48592054, 0.25186266, 0.93769755, 0.28565487, 0.47219067,
	0.99054882, 0.13155240, 0.47110470, 0.98556600, 0.84397623,
	0.12875246, 0.90953202, 0.49129015, 0.23792727, 0.79481194,
	0.44337770, 0.96564297, 0.67749118, 0.55684872, 0.27286897,
	0.79538393, 0.61965356, 0.22487929, 0.02226018, 0.49248200,
	0.42247006, 0.91797788, 0.99250134, 0.23449967, 0.52531508,
	0.10246337, 0.78685622, 0.34310922, 0.89892996, 0.40454552,
	0.68608407, 0.30752487, 0.83601319, 0.54956031, 0.63777550,
	0.82199797, 0.24890696, 0.48801123, 0.48661910, 0.51223987,
	0.32969635, 0.31075073, 0.21393155, 0.73453207, 0.15565705,
	0.58584522, 0.28976728, 0.97621478, 0.61498701, 0.23891470,
	0.28518540, 0.46809591, 0.18371914, 0.37597910, 0.13492176,
	0.66849449, 0.82811466, 0.56240330, 0.37548956, 0.27562998,
	0.27521910, 0.74096121, 0.77176757, 0.13748143, 0.99747138,
	0.92504502, 0.09175241, 0.21389176, 0.21766512, 0.31183245,
	0.23271221, 0.21207367, 0.57903312, 0.77523344, 0.13242613,
	0.31037988, 0.01204835, 0.71652949, 0.84487594, 0.14982178,
	0.57423142, 0.45677888, 0.48420169, 0.53465428, 0.52667473,
	0.46880526, 0.49849733, 0.05670710, 0.79022476, 0.03872047,
	0.21697212, 0.20443086, 0.28949326, 0.81678186, 0.87629474,
	0.92297064, 0.27373097, 0.84625273, 0.51505586, 0.00582792,
	0.33295971, 0.91848412, 0.92537226, 0.91760033, 0.07541125,
	0.71745848, 0.61158698, 0.00941650, 0.03135554, 0.71527471,
	0.24821915, 0.63636652, 0.86159918, 0.26450229, 0.60160194,
	0.35557725, 0.24477500, 0.07186456, 0.51757096, 0.62120362,
	0.97981062, 0.69954667, 0.21065616, 0.13382753, 0.27693186,
	0.59644095, 0.71500764, 0.04110751, 0.95730081, 0.91600724,
	0.47704678, 0.26183479, 0.34706971, 0.07545431, 0.29398385,
	0.93236070, 0.60486023, 0.48015011, 0.08870451, 0.45548581,
	0.91872718, 0.38142712, 0.10668643, 0.01397541, 0.04520355,
	0.93822273, 0.18011940, 0.57577277, 0.91427606, 0.30911399,
	0.95853475, 0.23611214, 0.69619891, 0.69601980, 0.76765372,
	0.58515930, 0.49479057, 0.11288752, 0.97187699, 0.32095365,
	0.57563608, 0.40760618, 0.78703383, 0.43261152, 0.90877651,
	0.84686346, 0.10599030, 0.72872803, 0.19315490, 0.66152912,
	0.10210518, 0.06257876, 0.47950688, 0.47062066, 0.72701157,
	0.48915116, 0.66110261, 0.60170685, 0.24516994, 0.12726050,
	0.03451185, 0.90864994, 0.83494878, 0.94800035, 0.91035206,
	0.14480751, 0.88458997, 0.53498312, 0.15963215, 0.55378627,
	0.35171349, 0.28719791, 0.09097957, 0.00667896, 0.32309622,
	0.87561479, 0.42534520, 0.91748977, 0.73908457, 0.41793223,
	0.99279792, 0.87908370, 0.28458072, 0.59132853, 0.98672190,
	0.28547393, 0.09452165, 0.89910674, 0.53681109, 0.37931425,
	0.62683489, 0.56609740, 0.24801549, 0.52948179, 0.98328855,
	0.66403523, 0.55523786, 0.75886666, 0.84784685, 0.86829981,
	0.71448906, 0.84670080, 0.43922919, 0.20771016, 0.64157936,
	0.25664246, 0.73055695, 0.86395782, 0.65852932, 0.99061803,
	0.40280575, 0.39146298, 0.07291005, 0.97200603, 0.20555729,
	0.59616495, 0.08138254, 0.45796388, 0.33681125, 0.33989127,
	0.18717090, 0.53545811, 0.60550838, 0.86520709, 0.34290701,
	0.72743276, 0.73023855, 0.34195926, 0.65019733, 0.02765254,
	0.72575740, 0.32709576, 0.03420866, 0.26061893, 0.56997511,
	0.28439072, 0.84422744, 0.77637570, 0.55982168, 0.06720327,
	0.58449067, 0.71657369, 0.15819609, 0.58042821, 0.07947911,
	0.40193792, 0.11376012, 0.88762938, 0.67532159, 0.71223735,
	0.27829114, 0.04806073, 0.21144026, 0.58830274, 0.04140071,
	0.43215628, 0.12952729, 0.94668759, 0.87391019, 0.98382450,
	0.27750768, 0.90849647, 0.90962737, 0.59269720, 0.96102026,
	0.49544979, 0.32007095, 0.62585546, 0.03119821, 0.85953001,
	0.22017528, 0.05834068, 0.80731217, 0.53799961, 0.74166948,
	0.77426600, 0.43938444, 0.54862081, 0.58575513, 0.15886492,
	0.73214332, 0.11649057, 0.77463977, 0.85788827, 0.17061997,
	0.66838056, 0.96076133, 0.07949296, 0.68521946, 0.89986254,
	0.05667410, 0.12741385, 0.83470977, 0.63969104, 0.46612929,
	0.10200126, 0.01194925, 0.10476340, 0.90285217, 0.31221221,
	0.32980614, 0.46041971, 0.52024973, 0.05425470, 0.28330912,
	0.60426543, 0.00598243, 0.97244013, 0.21135841, 0.78561597,
	0.78428734, 0.63422849, 0.32909934, 0.44771136, 0.27380750,
	0.14966697, 0.18156268, 0.65686758, 0.28726350, 0.97074787,
	0.63676171, 0.96649494, 0.24526295, 0.08297372, 0.54257548,
	0.03166785, 0.33735355, 0.15946671, 0.02102971, 0.46228045,
	0.11892296, 0.33408336, 0.29875681, 0.29847692, 0.73767569,
	0.02080745, 0.62980060, 0.08082293, 0.22993106, 0.25031439,
	0.87787525, 0.45150053, 0.13673441, 0.63407612, 0.97907688,
	0.52241942, 0.50580158, 0.06273902, 0.05270283, 0.77031811,
	0.05113352, 0.24393329, 0.75036441, 0.37436336, 0.22877652,
	0.59975358, 0.85707591, 0.88691457, 0.85547165, 0.36641027,
	0.58720133, 0.45462835, 0.09243817, 0.32981586, 0.07820411,
	0.25421519, 0.36004706, 0.60092307, 0.46192412, 0.36758683,
	0.98424170, 0.08019934, 0.68594024, 0.45826386, 0.29962317,
	0.79365413, 0.89231296, 0.49478547, 0.87645944, 0.23590734,
	0.28106737, 0.75026285, 0.08136314, 0.79582424, 0.76010628,
	0.82792971, 0.27947652, 0.72482861, 0.82191216, 0.46171689,
	0.79189752, 0.96043686, 0.51609668, 0.88995725, 0.28998963,
	0.55191845, 0.03934737, 0.83033700, 0.49553013, 0.98009549,
	0.19017594, 0.98347750, 0.33452066, 0.87144372, 0.72106301,
	0.71272114, 0.71465963, 0.88361677, 0.85571283, 0.73782329,
	0.20920458, 0.34855153, 0.46766817, 0.02780062, 0.74898344,
	0.03680650, 0.44866557, 0.77426312, 0.91025891, 0.25195236,
	0.87319953, 0.63265037, 0.25552148, 0.27422476, 0.95217406,
	0.39281839, 0.66441573, 0.09158900, 0.94515992, 0.07800798,
	0.02507888, 0.39901462, 0.17382573, 0.12141278, 0.85502334,
	0.19902911, 0.02160210, 0.44460522, 0.14688742, 0.68020336,
	0.71323733, 0.60922473, 0.95400380, 0.99611159, 0.90897777,
	0.41073520, 0.66206647, 0.32064685, 0.62805003, 0.50677209,
	0.52690101, 0.87473387, 0.73918362, 0.39826974, 0.43683919,
	0.80459118, 0.32422684, 0.01958019, 0.95319576, 0.98326137,
	0.83931735, 0.69060863, 0.33671416, 0.68062550, 0.65152380,
	0.33392969, 0.03451730, 0.95227244, 0.68200635, 0.85074171,
	0.64721009, 0.51234433, 0.73402047, 0.00969637, 0.93835057,
	0.80803854, 0.31485260, 0.20089527, 0.01323282, 0.59933780,
	0.31584602, 0.20209563, 0.33754800, 0.68604181, 0.24443049,
	0.19952227, 0.78162632, 0.10336988, 0.11360736, 0.23536740,
	0.23262256, 0.67803776, 0.48749791, 0.74658435, 0.92156640,
	0.56706407, 0.36683221, 0.99157136, 0.23421374, 0.45183767,
	0.91609720, 0.85573315, 0.37706276, 0.77042618, 0.30891908,
	0.40709595, 0.06944866, 0.61342849, 0.88817388, 0.58734506,
	0.98711323, 0.14744128, 0.63242656, 0.87704136, 0.68347125,
	0.84446569, 0.43265239, 0.25146321, 0.04130111, 0.34259839,
	0.92697368, 0.40878778, 0.56990338, 0.76204273, 0.19820348,
	0.66314909, 0.02482844, 0.06669207, 0.50205581, 0.26084093,
	0.65139159, 0.41650223, 0.09733904, 0.56344203, 0.62651696,
	0.67332139, 0.58037374, 0.47258086, 0.21010758, 0.05713135,
	0.89390629, 0.10781246, 0.32037450, 0.07628388, 0.34227964,
	0.42190597, 0.58201860, 0.77363549, 0.49595133, 0.86031236,
	0.83906769, 0.81098161, 0.26694195, 0.14215941, 0.88210306,
	0.53634237, 0.12090720, 0.82480459, 0.75930318, 0.31847147,
	0.92768077, 0.01037616, 0.56201727, 0.88107122, 0.35925856,
	0.85860762, 0.61109408, 0.70408301, 0.58434977, 0.92192494,
	0.62667915, 0.75988365, 0.06858761, 0.36156496, 0.58057195,
	0.13636150, 0.57719713, 0.59340255, 0.63530602, 0.22976282,
	0.71915530, 0.41162531, 0.63979565, 0.09931342, 0.79344045,
	0.10893790, 0.84450224, 0.23122236, 0.99485593, 0.73637397,
	0.17276368, 0.13357764, 0.74965804, 0.64991737, 0.61990341,
	0.41523170, 0.05878239, 0.05687301, 0.05497131, 0.42868366,
	0.42571090, 0.25810502, 0.89642955, 0.30439758, 0.39310223,
	0.11357431, 0.04288255, 0.23397550, 0.11200634, 0.85621396,
	0.89733974, 0.37508865, 0.42077265, 0.68597384, 0.72781399,
	0.19296476, 0.61699087, 0.31667128, 0.67756410, 0.00177323,
	0.05725176, 0.79474693, 0.18885238, 0.06724856, 0.68193156,
	0.42202167, 0.22082041, 0.28554673, 0.64995708, 0.87851940,
	0.29124547, 0.61009521, 0.87374537, 0.05743712, 0.69902994,
	0.81925115, 0.45653873, 0.37236821, 0.31118709, 0.52734307,
	0.39672836, 0.38185294, 0.30163915, 0.17374510, 0.04913278,
	0.90404879, 0.25742801, 0.58266467, 0.97663209, 0.79823377,
	0.36437958, 0.15206043, 0.26529938, 0.22690047, 0.05839021,
	0.84721160, 0.18622435, 0.37809403, 0.55706977, 0.49828704,
	0.47659049, 0.24289680, 0.88477595, 0.07807463, 0.56245739,
	0.73490635, 0.21099431, 0.13164942, 0.75840044, 0.66877037,
	0.28988183, 0.44046090, 0.24967434, 0.80048356, 0.26029740,
	0.30416821, 0.64151867, 0.52067892, 0.12880774, 0.85465381,
	0.02690525, 0.19149288, 0.49630295, 0.79682619, 0.43566145,
	0.00288078, 0.81484193, 0.03763639, 0.68529083, 0.01339574,
	0.38405386, 0.30537067, 0.22994703, 0.44000045, 0.27217985,
	0.53831243, 0.02870435, 0.86282045, 0.61831306, 0.09164956,
	0.25609707, 0.07445781, 0.72185784, 0.90058883, 0.30070608,
	0.94476583, 0.56822213, 0.21933909, 0.96772793, 0.80063440,
	0.26307906, 0.31183306, 0.16501252, 0.55436179, 0.68562285,
	0.23829083, 0.86511559, 0.57868991, 0.81888344, 0.20126869,
	0.93172350, 0.66028129, 0.21786948, 0.78515828, 0.10262106,
	0.35390326, 0.79303876, 0.63427924, 0.90479631, 0.31024934,
	0.60635447, 0.56198079, 0.63573813, 0.91854197, 0.99701497,
	0.83085849, 0.31692291, 0.01925964, 0.97446405, 0.98751283,
	0.60944293, 0.13751018, 0.69519957, 0.68956636, 0.56969015,
	0.46440193, 0.88341765, 0.36754434, 0.89223647, 0.39786427,
	0.85055280, 0.12749961, 0.79452122, 0.89449784, 0.14567830,
	0.45716830, 0.74822309, 0.28200437, 0.42546044, 0.17464886,
	0.68308746, 0.65496587, 0.52935411, 0.12736159, 0.61523955,
	0.81590528, 0.63107864, 0.39786553, 0.20102294, 0.53292914,
	0.75485590, 0.59847044, 0.32861691, 0.12125866, 0.58917183,
	0.07638293, 0.86845380, 0.29192617, 0.03989733, 0.52180460,
	0.32503407, 0.64071852, 0.69516575, 0.74254998, 0.54587026,
	0.48713246, 0.32920155, 0.08719954, 0.63497059, 0.54328459,
	0.64178757, 0.45583809, 0.70694291, 0.85212760, 0.86074305,
	0.33163422, 0.85739792, 0.59908488, 0.74566046, 0.72157152,
}

func TestMt(t *testing.T) {
	key := []uint32{0x123, 0x234, 0x345, 0x456}
	mt.InitByArray(key)
	for _, want := range expectedInt {
		have := mt.GenrandInt32()
		if have != want {
			t.Errorf("wrong output: %d != %d", have, want)
		}
	}
	for _, want := range expectedFloat {
		have := mt.GenrandReal2()
		if fmt.Sprintf("%.8f", have) != fmt.Sprintf("%.8f", want) {
			t.Errorf("wrong output: %.8f != %.8f", have, want)
		}
	}
}
