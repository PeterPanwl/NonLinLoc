Vel2Grid run/chichi.in
Grid2GMT run/chichi.in model/layer.P.mod gmt/ V G 1 0 1 301
#evince gmt/layer.P.mod.VG.ps
Grid2Time run/chichi.in
Grid2GMT run/chichi.in time/layer.P.AURF.time gmt/ V G 0 0 0 301
#evince gmt/layer.P.AURF.time.VG.ps
Grid2GMT run/chichi.in time/layer.P.AURF.angle gmt/ V G 0 0 0 301
#evince gmt/layer.P.AURF.angle.VG.ps
Time2EQ run/chichi.in
#more obs/synth.obs
NLLoc run/chichi.in
#more loc/chichi.19950421.080259.grid0.loc.hyp
Grid2GMT run/chichi.in loc/chichi.19950421.080259.grid0.loc gmt/ L S
#evince gmt/chichi.19950421.080259.grid0.loc.LS.ps
LocSum ./run/chichi 1 loc/chichi "loc/chichi.*.*.grid0.loc"
Grid2GMT run/chichi.in loc/chichi gmt/ L E101
#evince gmt/chichi.LE_101.ps
