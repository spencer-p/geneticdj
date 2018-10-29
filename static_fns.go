package main

var (
	interesting_fns = []string{
		"(/ (+ (* (/ t (& (i [-66 -57 11 -4 -77 -64 6 42 -21 4 -45 0] (- (^ (i [70 -3 -52 -20] -13) (^ t t)) (| (% t -58) (/ -93 -17)))) -88)) (| (/ 39 (+ 75 (| (+ (& 91 -53) (/ -43 -57)) (>> (% 30 -29) (- t -4))))) (>> (<< 68 -36) -87))) -15) (i [51 86 58 81 97 -9 13 -95 17] (| (<< (>> (+ (<< (/ 93 49) -72) (% 17 -27)) 36) (| 20 (+ (+ (<< (- 15 -23) (i [-58 -27 68] 36)) (>> (% 32 64) (| t -34))) -11))) -60)))",
		"(^ (& (>> (+ t (<< (/ (>> (>> (* 35 -78) -34) -15) (% 56 (/ (^ 21 34) (>> 58 -64)))) 60)) (+ (| (| 46 82) 37) (* (& (| (% (& -5 -86) (<< -44 22)) (/ (^ -74 -13) (>> 37 t))) (+ (% (i [-29 82 -22 97 40 -47 6 -87 89 28 -30 66 -50 66 46 100 -75 -48 -57 42 -66 -67 96 24 34 31] 46) -69) -12)) 56))) (<< -30 60)) (<< (% (* 5 (+ (& 20 (* (& t t) -74)) (+ t (i [64 -23 15 78 -3 -84 39 -68 31 -17 -77] (+ (& -65 -92) (>> -52 29)))))) (& t (% (* t (% (>> (+ 98 -82) -67) (& (% -21 t) -83))) (& t (<< (* t 63) 38))))) -10))",
		"(* (<< (<< (^ (i [62 -32 86 -88 -48 97 55 -12 -13 45 -31 -39 24] -22) (i [29 99 -43 97 -14 -60 3 -22 -86 -12 99 -27 -51 80 -83 38 70 -27 -42 -99 -57] (* (>> (| (- 70 33) (^ -5 -95)) (* (+ t -67) (- t 37))) (% (i [37 64 -11 74 -77 -25 -79 60 -61] (>> t t)) -50)))) (| (i [-24 25 -6 4 -72 -100 -15 33 -53 -29 -77 8 36 42 8 31 76 -46 -88 58 -72 -79 -77 -57 -100 -46 27 -43 90 62] (- (<< 74 (+ (| t 41) (/ 55 26))) (^ (% (% -16 -62) (% 73 -26)) (- (% 26 -66) (>> t t))))) -95)) (^ (+ (* (^ (- (| 70 t) (>> (+ -27 -26) (i [87 -66 -26 100 -18 86 14 -93 -27 77 -80 47 -58 78] 90))) (& t -26)) -68) t) (/ -27 (- t 34)))) (% (^ (* (| (& (>> (* 26 39) t) 39) (| (& (i [-98 85 -64 55 -88 -100 -30 -25 23 38 72 98 -42 -37 21 49 -20 -48 -5 -62 -71 31 28 -42 27 0] (& -73 19)) -74) (>> (<< (& 99 t) (% 65 -76)) (<< -87 -13)))) (i [-66 -37 -50 97 -81 -3 -57 -74 72 92 -83 -61 68 93 -42 54 14 15 99 -91 -87 84 -18 -82 -9 90 -94 81] (* (+ (* -60 (>> 62 84)) (% t -66)) (& (^ 30 (^ 68 -83)) 47)))) 70) (i [-38 76 62 -53 87 11 0 -36 4 10 44 -55 -5 93 57 40 87 -83 83 54 50 65 39 70 -6 -95 -38 56] (/ (* (^ (& (+ (/ 92 99) -8) (| t (>> -59 t))) (& -90 (- 91 (| t 54)))) (* (i [86 -88 54 45 -8 65 42 48 96 57 -52 8 65 16 -64 97 27 88 -15 -13 8] (i [82 -12 36 -37] (i [-75 96 -22 -62 -20 26 -73 25 -10 66 72 -91 -64 44 -32] 69))) -29)) (>> (* t (& 25 17)) 72)))))",
		"(| (^ (& (^ 3 8) (^ 7 0)) (/ (^ 1 t) (i [6 8 8] 2))) t)",
		"(^ (| (- (/ (% (| (+ (i [10 4 7] (i [7 8 7] t)) (/ (<< t 5) (>> 5 t))) 3) (<< t t)) (& (| (- (+ 4 (/ t 8)) (% (/ t 1) (% 4 t))) (* (/ (i [3 10 7 5 5 9 9] t) (<< t t)) t)) (+ t (* (- (^ t t) t) (* (% 7 t) (/ t t)))))) (<< (/ (>> 6 (| (i [6 1 4 5 6] (>> 10 5)) (% 4 t))) (^ 7 9)) 6)) (% (+ t (/ t t)) (+ (<< (i [2 9 4 2] (^ t (* (i [5 10 6 9 2] t) (/ 5 t)))) 9) t))) (| (^ (^ (^ (>> (% (i [9 5 8 4] (* t t)) (| (/ 6 1) (% t 1))) (/ (/ (* 10 5) (- t 1)) (+ (| 5 5) (i [5 1 2 3 4 2 9] t)))) t) (| 2 (% (^ (<< (& t t) 6) 8) 2))) (& (- (^ (<< (- (% t 1) 3) (| (* t t) (- 8 1))) t) (& t (* (^ t 2) 10))) (| t (i [9 2 6 1 7] (| (- (<< 7 t) (* 10 1)) t))))) (- (<< (| (| (* (% t (* 2 t)) (% (i [10 2 8 6] t) (>> t t))) (>> (>> (& 5 t) (& 7 9)) (>> (| t t) (* t 2)))) t) (| (+ 4 (| (>> (| 5 2) 10) (/ (>> 4 10) (>> 4 8)))) (+ (| (% (& 1 t) t) (>> (| 3 9) (| 5 t))) t))) (^ 8 (i [1 4 5 7] 8)))))",
		"(| (& t (& t (<< (- (* (% (i [7 2 2 9 5] (i [10 10 4 10 2] 6)) (| (% 5 t) (+ t 5))) (i [7 6 1 2 3] t)) 7) (/ (<< (* t (- (& t t) t)) (% (+ (* 5 t) (% t 7)) 10)) (/ 5 (+ (/ (% t 5) t) (% (- t t) (% 10 t)))))))) (- 6 (- (i [9 8 9 6 6 4] (^ 7 9)) (^ (| (/ (| (* (% 4 t) (| 9 t)) (i [2 3 3 10] 2)) 7) (^ (<< (& 1 (>> t t)) (& (% 5 t) (| 2 t))) (/ t (+ (- t 9) t)))) (^ (^ (* (* (/ 2 4) (| 8 2)) t) (>> (* (/ t t) (i [6 2 9 4 8] t)) (& (| t t) (& 4 3)))) (* t (<< 5 (* (+ t t) t))))))))",
		"(& (& (i [10 6 6 7] (i [6 8 7 9 7 10] (+ (>> 7 t) (* (% (i [3 2 3 9 4 5 5] (i [10 10 2] 5)) t) (* 9 (/ (+ t t) (i [3 2 7 1 9 9 3] 6))))))) (^ (/ (| (i [2 2 2 5] (+ (/ (- t t) (- t 7)) (| (- t 8) (% 5 t)))) (| t 9)) (% (i [5 10 6 3 1] (>> (* 6 (- 4 10)) (% 5 (<< t t)))) (<< (i [9 10 8] (* (i [5 4 5 1 9 5 1] 10) (% t 9))) (- (* (- 8 t) (^ 9 t)) (i [4 7 5 6 8 8] (^ 7 3)))))) (/ (* (% 9 10) (- (| (/ (- 7 4) (>> 10 t)) (- (- t t) (& 7 t))) (& t t))) (/ (<< (- t (<< (| 3 t) t)) (>> (>> (| 1 t) (i [8 7 2 6 3 6 7] t)) (^ (<< 6 t) 10))) t)))) (* (^ 5 (>> (^ (<< 2 6) (/ (* (<< t (| 2 t)) (- (i [7 4 5 2 8 10 9] 3) (<< t t))) (% (| t (- t t)) (+ t (>> t t))))) (i [2 7 3 2 3 5] (<< t (% (| t (% 1 t)) (* (+ t 3) (i [1 2 9] t))))))) (i [1 4 1] (^ t (^ (* (>> 6 (| (& t 10) (^ 5 6))) (| (i [5 1 4 1] (>> 3 t)) (<< (/ t 5) (& 2 10)))) (+ (- (<< 7 (/ 3 t)) (% (>> t 10) (& t 2))) (^ (- t (& 4 t)) (<< t (- t t)))))))))",
		"(^ (i [4 2 6] (% t (>> (| (| (/ (i [2 4 3 7 4] t) (>> (>> 1 5) 1)) (* 3 (& (/ t 8) (>> 3 8)))) (& (/ (+ t t) (+ (* t t) (i [5 1 4 9 5 10 2] t))) t)) (+ (<< (i [8 1 5 2 6 6] (i [8 2 9 10 6] t)) (^ 8 (| t (i [8 4 7 5 5 7 10] t)))) (+ (/ (& (/ t 8) (+ 7 t)) (i [1 1 5 2 4 1] (<< t t))) (<< (- 5 (| t t)) (^ (^ 1 t) (- 10 6)))))))) (- (& 2 (i [8 5 9 5 10 3] t)) (/ t (- (+ (^ (+ 4 (| (/ t t) (| t 6))) (+ (* (& t 7) (| t t)) (>> (>> 8 8) (i [2 3 4] 4)))) 9) (^ (- (i [8 6 1 2 8 7] (<< 1 (% 8 t))) (+ (* (<< 10 t) t) (+ (i [3 7 7 1 3 7 6] 8) (i [2 5 1 1 3 7] t)))) t)))))",
		"(- (+ (+ (i [2 8 9 9] t) (% (* (& 8 (* (- 1 t) t)) t) (>> (* (| 4 (% 7 (- 1 t))) (+ (% (>> 8 6) (^ 2 2)) (<< (| t t) (& t 3)))) (% (+ (<< t (<< t t)) (& (+ 4 10) (* 10 3))) (| t (>> (* 6 9) t)))))) (<< (- (% (/ 6 (>> (>> (>> 2 10) t) (+ (^ 4 t) (% t 3)))) (| t t)) (- (& 5 t) 10)) (| (^ (& t t) (- t (& (/ (i [10 4 5 9 4 9] t) (+ 8 8)) (<< (>> 2 9) t)))) t))) (<< (<< (* (i [6 1 1 4 4] (>> (>> (| (i [10 7 8 7 4 6 9] t) (>> t 7)) (i [4 1 8 6 5 7 3] (^ 8 t))) (<< (<< (* t 4) (& 8 t)) 3))) t) (| t 8)) (i [5 2 9 1 10 4 1] (^ (+ (i [7 6 3 5 6 4 1] (i [7 2 9 10] 7)) (^ (* (+ (>> t t) t) (<< (* 7 4) t)) (& (>> (| t 10) (/ t 10)) (- (& 6 6) (>> t t))))) (& (/ (i [1 9 1 4 8 5 7] (<< (/ t 3) (& 9 1))) (- 10 (- 2 (% 6 t)))) (i [9 8 9 3 9] t))))))",
		"(/ (<< (* (+ (^ (+ (% (/ (/ 9 t) (% t t)) (+ (^ t 3) (<< t 1))) (/ t (/ (& t t) (- t 6)))) (| (<< 8 (| (^ t t) (& t t))) (% (% (>> 4 8) (* t 8)) (>> (i [3 6 9 2] 6) (- 6 9))))) (| (* (| (- (+ 8 t) 4) t) 9) (- t (>> (^ (/ 9 t) (i [4 8 8] t)) (^ 5 (^ 9 5)))))) (^ 4 10)) (* (| (^ (^ (* (% (^ t 9) (i [8 6 6 9] 1)) (>> t (+ 2 10))) (* (* (- 3 t) (>> t 1)) (* t (+ t 9)))) (- (^ (>> (+ 4 5) (- t t)) (^ (<< t t) (<< t t))) (- (* 10 (- t t)) (i [7 2 9 10 7 3] (<< t t))))) (* (>> (>> 4 9) (& 9 (<< (^ t t) (% 4 4)))) (* (i [3 10 5 10 6 8] (* (* t t) (| t 9))) t))) t)) (& (+ (& 2 (<< (- (* (^ (| 9 1) t) (- (>> t t) (* 1 10))) (- (% t (i [5 3 3 10] t)) (% (& 2 9) (& t 7)))) (% t (% (& 9 (^ t t)) (+ t (| 4 t)))))) (& (i [6 6 3 5] (>> t 1)) t)) (<< (% t (* t (^ 6 (<< 1 (& t (- t t)))))) t)))",
		"(& (<< (^ (- (% (i [4 7 1 8 1] (^ (<< 8 (<< 1 t)) (/ t (- t 6)))) (| (& (& (- 5 t) (- 10 1)) 6) (/ t t))) (<< (| (/ (% 9 t) (& t (i [4 10 6 8] 10))) (^ t (- (& t t) t))) (i [7 3 4 5 3 7 7] (<< 1 (/ (>> t 7) 9))))) (| 9 (| (i [9 6 1 6] (/ (i [5 1 2 6 4 4] (| t t)) (^ t (^ t 10)))) 1))) 1) (/ (| (+ (* (/ (<< (/ t (& t t)) 8) (& (<< (& t t) (/ 3 t)) t)) (<< (* (| (^ t t) (+ 8 t)) t) t)) (^ (- (<< 2 (>> t (% 9 1))) (/ (^ 9 8) (i [4 8 9 7] (i [7 8 10 5 10 9 10] 4)))) t)) (<< (i [2 8 8 4 7] (i [9 4 3 1 1] (^ 7 (- (i [5 1 8 4 5 5 2] t) (| t t))))) (- (& (+ (^ t 5) 9) (* t (| t t))) t))) (/ (- 8 (/ 3 (/ (^ (^ (^ t 3) (/ 3 4)) 4) (>> 2 (i [3 7 10 3 4] (| 7 t)))))) (% t (% (i [9 1 2 1] (<< 6 (^ (| 8 t) (^ 9 t)))) (i [10 4 1 4 3 2 1] (* (/ (<< 5 9) 3) 4)))))))",
		"(& t (^ (>> (^ (% (i [4 7 7 5 6 7 10] 2) (* t (^ (* 7 t) (/ (* t 1) (/ 4 9))))) 5) (* (- (i [10 1 3 10 8 10] (i [6 1 8] (<< (% 1 8) (^ 10 t)))) (^ 3 (- 6 (/ (>> 3 3) (/ t t))))) (% t 3))) (+ (| (% (| (>> (>> (/ t 2) (* 10 t)) (/ (^ t 8) (% 2 6))) (>> (/ (i [9 1 7 4 1] t) t) (<< (% 2 9) (& t t)))) (^ (<< (>> (& 5 t) 9) (^ 9 7)) (& 7 (- (/ 5 9) (- t 1))))) (| (- (& t 1) (& (<< 1 (>> 2 9)) (/ t (/ t 8)))) (^ (>> (/ (| t t) (% t 9)) (& (/ t 8) (& t t))) 5))) 3)))",
		"(/ (+ (>> (>> t (^ t 1)) (& (/ (+ t (% (& (- t t) 7) (/ (% 8 7) 4))) (* t (>> (| (<< t t) 5) t))) (+ (^ (& (- (/ 10 5) (^ 10 3)) (* (* 10 3) (- 9 5))) (/ (^ (& 7 6) (i [5 5 1 2 2 2] t)) (- (>> 8 5) (^ t t)))) (| (i [4 2 6 6 2] t) (<< (& (i [2 3 10 3 1 8 8] 3) (- 7 2)) 8))))) (% (- (| t (% (^ (& 8 (^ 9 t)) (<< (<< 7 t) (^ 1 t))) (- (+ 5 (+ t t)) 8))) (% (| (- (- 7 t) (& 1 (^ 2 6))) (* (| (i [6 1 2 6 4 9 1] 9) (/ 10 t)) (^ (* t 7) (& t 1)))) (>> (* (>> (| 3 5) (| t 10)) (& (i [1 8 9 7 2] 6) (<< t t))) 6))) (| (- 10 (>> (| (<< 7 (& t t)) t) (i [6 10 9 4 9 7] (^ (% 1 t) (& t t))))) (- (| 10 (- 9 (% (+ 6 t) (i [5 6 5 7] t)))) (| (| (+ (& t t) t) (>> (i [8 7 10 10 5 9] 2) (- t 2))) 9))))) (/ (/ (^ (>> (| (i [10 5 8 4 6 1] (& (i [4 4 8 9] 4) (>> t 10))) (% (i [9 3 2 4 6] (| 5 1)) (>> (i [6 7 7 6 5] 5) (% t t)))) (^ (<< t (* (^ t t) (<< t t))) (* (* t (i [9 8 6 2 6 6 10] 5)) (^ 8 (<< t 3))))) (^ (>> (^ (- (% 6 5) 5) (| (% 3 10) (* 8 5))) t) (- (<< (% (& t t) (>> t 10)) t) (>> (^ (i [4 7 1 3] 5) (| 5 9)) (i [2 9 1 4 8 8 8] t))))) (>> (^ 3 (i [2 10 3] (^ 2 (<< (>> 2 9) t)))) (| 6 (i [10 6 4 10] (<< t (^ (& t t) (- 10 t))))))) (i [6 8 9 1 10 6 3] 6)))",
		"(* (& (<< t (+ (- (+ (% (& (>> 3 t) (% t t)) (% (& 9 6) (/ t 9))) (| (+ (/ t t) (% t 6)) (>> (>> t t) (| t t)))) (% (+ 7 (/ (% 10 5) (% 10 6))) 3)) (<< (- (>> (+ 3 (% t t)) (| (+ 6 8) (^ t t))) 5) (+ (<< (& 8 (+ 8 5)) (& t (^ 9 t))) (<< (i [3 7 2 6] (- 1 10)) (+ (>> t 2) (& t t))))))) (% (& 8 (- (/ (i [10 5 9 2 4 10 1] (% (>> t 9) (- 2 t))) (& (+ (& t 2) (^ t t)) (+ (+ 3 6) (* t 2)))) (^ (i [8 10 7 6 10 5] (^ t (i [5 10 9 7 9] t))) (>> (| (^ 6 t) (- 4 5)) (/ (| t t) (+ 3 6)))))) t)) (- (/ (>> (i [10 1 7 4] 4) (>> (& (- (| (+ 3 2) t) 4) 5) t)) (/ (>> (i [7 10 6 9 10 7 10] (<< t (^ (+ t 4) 1))) (/ 7 (* (/ (+ 9 6) (^ 9 t)) (& (| t 5) 10)))) (+ (i [5 3 7 9 7] (^ t (+ (/ t 10) t))) 8))) t))",
		"(* 3 (>> (- (/ (<< t (& (- (% (^ 6 t) (i [10 9 5 1 3 9] t)) (<< 10 (<< t 3))) (% (* 7 (^ 3 10)) (>> (i [5 4 3 7 8 3] t) (- t t))))) 4) (<< (- 6 t) (- (| t 7) (>> (* (i [10 3 3 2 6 6 2] (| 5 t)) (/ (>> t t) (/ t t))) (* (<< (<< t t) (% t 10)) (- t (% 4 t))))))) (+ 1 (/ 7 (* t (>> (/ (i [1 8 6 10 2] (i [3 4 3] 3)) t) (>> (* (% 9 t) (+ 9 t)) (+ (* 5 6) t))))))))",
		"(- (<< (+ (* t 2) (i [1 3 2] (& (- t (| (% (^ t 1) (& t t)) (| 10 t))) (>> (& (>> (- 4 9) (i [4 4 6 4] t)) (| t (>> t 6))) (^ 10 (| 4 (% 9 t))))))) (| (| t t) (- (/ (/ (<< (>> (/ 8 9) (^ t t)) (i [1 5 2 9 3 1 1] t)) (- (i [6 8 5] 7) (% (% 4 t) 9))) (* (% (+ t t) (>> 10 (/ t t))) (>> (>> (^ 4 8) (^ 10 t)) (| t (i [2 1 8] t))))) (>> (- (<< (% (i [7 9 2 8 10 7 10] 7) (- t 4)) (/ (i [2 10 8] t) (+ t 8))) (% t (>> (| t t) 8))) (* (+ (| (+ t 7) (/ 10 1)) (* (| t t) (+ 4 4))) (% (<< (- t t) (- t t)) 7)))))) (& t (>> (% (* (/ (<< 7 (* (<< 5 10) t)) (| (& t (<< t 1)) (& (* 7 t) t))) (/ 3 (/ 4 (- (| t 1) (>> 10 7))))) 5) (+ (^ 9 (^ 10 (>> (+ (& t t) (| t 8)) (<< 10 t)))) 9))))",
		"(^ (>> (* t (- (^ (& (& 2 (/ (+ t 5) (<< t t))) t) (% (| t 1) (& t (<< (^ t 9) (* t t))))) (* (/ (| (& (<< t t) (+ 5 t)) 4) t) (i [3 3 4] (* t (+ (>> 7 7) (| t 8))))))) (i [10 8 7 8 10] (& (+ (+ (<< 2 (* t (& 5 t))) t) (i [7 1 5 6 5] (* t t))) 8))) (/ (& (% (+ (/ t 8) 6) t) 1) (+ t t)))",
		"(- 6 (& (- (| t (i [5 7 10 6 4] (<< (* t (/ (/ t t) (* 3 1))) (| (% (>> 2 t) (^ t 1)) (- (- t 1) (* t t)))))) (| (>> (^ (<< (| (& 5 t) (| t 3)) (>> 4 (- t 5))) 6) (* (^ (| (^ 7 8) 3) (i [6 6 2 1 2] (<< t t))) (>> (- (+ 10 t) (% t 6)) t))) (+ (+ (% (% (^ 7 t) 2) (i [7 1 5 1 10 8] (& 10 8))) t) (+ (+ (>> (+ 9 7) (- 7 t)) (- (* 9 9) (/ t 8))) t)))) (/ t (& (- (- (>> t (/ (^ 6 2) (<< 10 t))) t) (i [3 6 8 3 2] (% (+ (/ t t) (>> t 2)) (& (- 1 8) (^ t t))))) t))))",
		"(+ (& (^ t (<< (^ (% (& (i [2 6 5 6 10 4 1] (| t 5)) (* (% t t) (+ 10 3))) t) (| (& (| (* 1 t) 1) (i [1 1 10 1 1 8] (>> t 9))) t)) (* (- (- 9 (& 2 (+ 7 1))) t) (- t (/ (% (& 2 t) t) (^ (/ t 2) (& t t))))))) 8) (% (<< (i [2 3 9 10 4 3] (| t (+ (<< 8 6) (>> t (- (+ t 9) t))))) (% 10 (>> (| 4 (i [10 6 6 10] (i [10 8 6 5 4 6 3] t))) (>> (* t (& (<< t 4) (/ 8 t))) (* (- 5 8) 8))))) (/ (/ (% (>> t (- (* (^ 10 t) 1) (/ (<< 10 t) t))) (^ 4 (i [6 2 5 6 10 10] (| (- 2 t) (% 1 7))))) (>> (* (^ (i [6 3 5] (- 7 8)) (* (* 8 t) (% 2 3))) 8) (<< 10 (>> 9 (- (>> t t) (>> 7 t)))))) (% (<< (& (/ (- (>> 3 t) (<< t 9)) t) (| (^ (<< 1 1) t) 7)) (& 8 (^ (* (>> t t) 5) (i [10 2 7 6 7] (% 9 t))))) (+ (^ (i [6 10 4 9] (^ (+ t 7) (^ t t))) (/ (& (& 6 4) (>> t t)) (& (>> 7 t) (/ t 8)))) (/ (<< (% t t) (& (+ 8 t) 9)) (/ t (| (& 3 1) 8))))))))",
	}
)