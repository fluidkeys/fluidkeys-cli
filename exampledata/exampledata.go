package exampledata

import (
	"github.com/fluidkeys/fluidkeys/fingerprint"
)

const ExamplePublicKey2 = `
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: pub   rsa1024/0x343CC240D350C30C 2018-09-10 [SC] [expires: 2038-09-07]
Comment: Key fingerprint = 5C78 E71F 6FEF B558 2965  4CC5 343C C240 D350 C30C
Comment: uid       <test2@example.com>
Comment: sub   rsa1024/0xA810C52C47D52528 2018-09-10 [E] [expires: 2038-09-07]

mI0EW5ZG7QEEANORDqZe+BRgs1YkGtjyVniNPMHHKQ1NixfygxBONvVyPU/2wE6P
xUzOlZz1yBs+c5iU9NOm9OSdcwRx0tU2IdBiNZYiaIM74l/6pgK4+v074AYaf/IZ
CBdwm6QDPVlUtXMztWd6j9T7mUNo/gKPg3TDJQ6tfQUHEKBz94pCOBbrABEBAAG0
Ezx0ZXN0MkBleGFtcGxlLmNvbT6IvwQTAQgAKQIbAwIZARYhBFx45x9v77VYKWVM
xTQ8wkDTUMMMBQJbmOOVBQklmqKoAAoJEDQ8wkDTUMMMHr0D/jHcl4cnNP5MvCg5
X/a7IM5NLPDZaGpMo4qD0OyFQ+96mE7A1rAHhYnUZ3iuU6+T0HKykRhihC0LspDh
4w6LBZVVpuVdHxRmPy59WxCoWB2ilog9oWCQWSJAx+zmuttAoms+25pANLBhr6T2
b9tY8wKkosByGXdCUdQymXnX868TuI0EW5ZG7QEEALsW6plBPn73//YoQ9LW1LGx
T/8XGYqlb6jUNqtboKxKsfG2/kewKfzUQnofyWUSo52VDE7xxKBm9+y6WAHyIfaR
OJUpLdwv0iI65iqqal0bj/wbDHFxEzzzC44i+J5n7SpvUi2xgdXSx/oho+HEWW3Z
qqJVmCPsUkS6OkrQn0G1ABEBAAGIvAQYAQgAJgIbDBYhBFx45x9v77VYKWVMxTQ8
wkDTUMMMBQJbmOOeBQklmqKxAAoJEDQ8wkDTUMMM2/4D/32A8qyedAyqWjzGa0MZ
MzbHj9vsKs+cOEYJZYwQzEt2sygldHSRWILeVhJfiKeaGGxw2ng+QYlwO2xlMo7H
063gpGk+EOpvg2ybhCa/Fd+sEhr6jzLulhRrGMkwmUiOMi1Wn1JHT3egaWf/Cvzp
qEMacEL0UbB7Unk9YLNqbnb5
=SdGd
-----END PGP PUBLIC KEY BLOCK-----
`

const ExamplePrivateKey2 = `
-----BEGIN PGP PRIVATE KEY BLOCK-----
Comment: encrypted with password "test2"

lQIGBFuWRu0BBADTkQ6mXvgUYLNWJBrY8lZ4jTzBxykNTYsX8oMQTjb1cj1P9sBO
j8VMzpWc9cgbPnOYlPTTpvTknXMEcdLVNiHQYjWWImiDO+Jf+qYCuPr9O+AGGn/y
GQgXcJukAz1ZVLVzM7Vneo/U+5lDaP4Cj4N0wyUOrX0FBxCgc/eKQjgW6wARAQAB
/gcDAk1l+s53NZj5/7FyAVqJj4kU0okUGTEloWwBml9mbOPg/5oGgIlz73Wbq8Bs
XAP3wdWthBcNmVRO9/BXMaffiosaQqd4O/lFwk4xAtUC2SZGIJ2a++nAwyeV9Vig
63kASRjdTJgu9n81AEJ8tp3XWc+/ntzFdVuWbUxlUuX1dTk4FOkLjJu3O78vzbW/
LIGiyeNDZIC5z6Zjd5vlqN9D6QwgGAr1BXhu3XyrBz0XWaCMEtNpmb3rqhzXL/kZ
qlCsF3xYhmJgOfoSarmsz+pKEZO2f7HAnE9BTX38GNZbkYj4SJDX8Yy0xeSZZvAM
0EPyZPiOTxf7r/V0LZgrFgk0NvLQqlT39Bz2zh/Cf2UMcX8J2vcUtJkxK6U9DkJk
MnxF6qLL7vWRcvobcY19BJf01kTnPQOMyQ2KksKj8Q5gtmcTMqCaFdPbBacexXnC
2G8m5mTURgdCcI2LPLhH56iOOelrwbsaX3zlPh/3urAOspZk4lNtjAu0Ezx0ZXN0
MkBleGFtcGxlLmNvbT6IvwQTAQgAKQIbAwIZARYhBFx45x9v77VYKWVMxTQ8wkDT
UMMMBQJbmOOVBQklmqKoAAoJEDQ8wkDTUMMMHr0D/jHcl4cnNP5MvCg5X/a7IM5N
LPDZaGpMo4qD0OyFQ+96mE7A1rAHhYnUZ3iuU6+T0HKykRhihC0LspDh4w6LBZVV
puVdHxRmPy59WxCoWB2ilog9oWCQWSJAx+zmuttAoms+25pANLBhr6T2b9tY8wKk
osByGXdCUdQymXnX868TnQIGBFuWRu0BBAC7FuqZQT5+9//2KEPS1tSxsU//FxmK
pW+o1DarW6CsSrHxtv5HsCn81EJ6H8llEqOdlQxO8cSgZvfsulgB8iH2kTiVKS3c
L9IiOuYqqmpdG4/8GwxxcRM88wuOIvieZ+0qb1ItsYHV0sf6IaPhxFlt2aqiVZgj
7FJEujpK0J9BtQARAQAB/gcDAu5o91JoURzr/4OU3wBWUhQIqB+/vbBH0WhCFLea
fj4ccqBaGPOeg0FZiAAmaKTkz//04DvlwwZHrw8jSGBYjauG2xp6BmtALvUphI8r
doQ3Hp4nkSr4DeSNFx33/1Mz5UV7UoD07PvEaU6hRiQ+c3CCKIcKJpK12E1P6EHH
zBzC6nUr8M94djQv7MksxvWMXhK4RstjPqBeC9t1TEy8olNjYmYV0amPNm6GHPid
BT83zftjX2syMpk1R3olYWl6PBr8R94AkgVVbPIjgPAYMLJf5zIdXh/qE3V38R33
Bek5zOdObG8ivJfY/xbaCgaYwu8uwLLGaM0Kc5NYOO7+e1bBjsOSvH+IzN0rjKwd
C9LdIO7w9Ge3uY9PGuasfmwyN8oZCVZkBx9Mz6trgKx/Dj02MKa0LzAOuJB349Jl
9HmdthDavetpBiy4qmlX2Xi4b4pBmJr/PLLcaubv3gEPHRaJIiX4g2EyVBP4s179
CtUV3UrJZliIvAQYAQgAJgIbDBYhBFx45x9v77VYKWVMxTQ8wkDTUMMMBQJbmOOe
BQklmqKxAAoJEDQ8wkDTUMMM2/4D/32A8qyedAyqWjzGa0MZMzbHj9vsKs+cOEYJ
ZYwQzEt2sygldHSRWILeVhJfiKeaGGxw2ng+QYlwO2xlMo7H063gpGk+EOpvg2yb
hCa/Fd+sEhr6jzLulhRrGMkwmUiOMi1Wn1JHT3egaWf/CvzpqEMacEL0UbB7Unk9
YLNqbnb5
=b0wI
-----END PGP PRIVATE KEY BLOCK-----
`

var ExampleFingerprint2 = fingerprint.MustParse("5C78 E71F 6FEF B558 2965  4CC5 343C C240 D350 C30C")

const ExamplePublicKey3 = `
-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: gpg --export-options export-minimal --armor --export
Comment: rsa1024/0x719BD63EF03BDC20 2018-09-10 [SC] [expires: 2038-09-07]
Comment: Key fingerprint = 7C18 DE4D E478 1356 8B24  3AC8 719B D63E F03B DC20
Comment: uid       <test3@example.com>
Comment: uid       Example Name <another@example.com>
Comment: uid       unbracketedemail@example.com
Comment: sub   rsa1024/0x409F66EB6D1336A7 2018-09-10 [E] [expires: 2038-09-07]

mI0EW5ZHDAEEAKwGAzK/KryNS+AZZElJxe4VRsRx1gYvGVCKyZrdqeFcovwmhDj9
IBQmNYD3+nJ/kyh/jDDFiYExrbWSXFJ7cZ7jpUNAyO0M/Rc0M/05RaOumRiWagcK
J5Wmx6Tco17NqlMbQr7i8ocKMi7bvZK0C75vVIo8fgk2Vmfq66KxMg6LABEBAAG0
Ezx0ZXN0M0BleGFtcGxlLmNvbT6IvwQTAQgAKQIbAwIZARYhBHwY3k3keBNWiyQ6
yHGb1j7wO9wgBQJbmNa/BQklmpWzAAoJEHGb1j7wO9wg9L8D/0CGkC6ySfr3kqM+
3EcXmI13z8O6fsBtDJAng/TVzH0PQ3RZOytoDzjI1Yq7LOoC+CmKY6ovodJPdckY
yy0uqdaLspZcuoqBHrosNAU8t8dKaoPi3jQf9kGMKoeuJnyTkoxMSMGxMGb1CDJc
0WT91W06D5tA9lruznjkPg6ZuRMctCJFeGFtcGxlIE5hbWUgPGFub3RoZXJAZXhh
bXBsZS5jb20+iNQEEwEKAD4CGwMFCwkIBwIGFQoJCAsCBBYCAwECHgECF4AWIQR8
GN5N5HgTVoskOshxm9Y+8DvcIAUCW5jWxAUJJZqVswAKCRBxm9Y+8DvcIBjRA/9B
Qv5qabfS2lbypANfTEaHW5GBwrxNuqJ4/yQwmPkD4l0kMGp/v6ZLR8AFP8J80zlb
HkXpssTtjRXeepCVS3d6YCIE/afBiliBLUsZuBn/iGWhHznH4QDY0W/7TR+H3DWm
zR0BuQzgLH/V2Ryi6ayoSIgK4AUgcb8ThNhk/obHL7QcdW5icmFja2V0ZWRlbWFp
bEBleGFtcGxlLmNvbYjUBBMBCgA+FiEEfBjeTeR4E1aLJDrIcZvWPvA73CAFAluY
3oUCGwMFCSWalbMFCwkIBwIGFQoJCAsCBBYCAwECHgECF4AACgkQcZvWPvA73CCB
rAQAl1ykuKAGwrCMyKYcHzf8lrDv8e64Lwxst4/LVu+oNNujT0BCFcZwv94NSLoz
2+M3OUEa/kAIMlyxxuijLXk6Zy4CVebTXNZn+a81JsZyKJe5joCimTm/5A0TpOaS
zVHRfgsipmuNPxmPFhn/wlEpbXr+o+8UOvfzTUp4D8rf/FC4jQRblkcMAQQA3exG
0sRizhHj3WzW6DHIXwlXgb4+1YRD5JGGPjlO5/E10mDZGUCyYNRY+5do7TlUZiVF
qHms1zhjdBk6o/lUPKFqB8SA5tqBQ2vHlBdOQAnT0/DvRcZ2V43LzgR5JEUhhttJ
+Qou12nwG92zuCuzv3A3Y/zcGJUQ2P7FM4JlkTEAEQEAAYi8BBgBCAAmAhsMFiEE
fBjeTeR4E1aLJDrIcZvWPvA73CAFAluY1s4FCSWalcIACgkQcZvWPvA73CCoRgP/
aLck/jKZV/+pvVI8dsYXoxMfTohPCF1USCbgGldY0EgPw/IURSzf696xCQ4cR9Rm
lnbiJG2uhesaqF7LXqwK/ZVR04sSSC7LMHSS5po2uWa++4BdqTd8a99u52fHYDFt
jLRaY5y3Qyhj30lmlTJc1Q/4kFVyIVBlHI/MUn7VdiM=
=2OVx
-----END PGP PUBLIC KEY BLOCK-----
`

const ExamplePrivateKey3 = `
-----BEGIN PGP PRIVATE KEY BLOCK-----
Comment: encrypted with password "test3"

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
D/iQVXIhUGUcj8xSftV2Iw==
=ZbqB
-----END PGP PRIVATE KEY BLOCK-----
`

var ExampleFingerprint3 = fingerprint.MustParse("7C18 DE4D E478 1356 8B24  3AC8 719B D63E F03B DC20")
