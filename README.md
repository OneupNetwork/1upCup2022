## 1up Cup 2022
![](https://img.shields.io/badge/date-2022--07--29-0385B1.svg)

![](/final_standings.jpg?raw=true)

### Problems

#### [1.Chocobo - Where's the Bishop?](https://codeforces.com/problemset/problem/1692/C)
 - 範例：練習各語言 stdio，二維陣列判斷左上右上左下右下位置的字元是否為 # 即可。

---

#### [2.Shiva - Odd/Even Increments](https://codeforces.com/problemset/problem/1669/C)
 - 範例：判斷奇/偶數位置的數字是否有相同的奇偶性即可。

---

#### [3.Ifrit - Kill the Monster](https://codeforces.com/problemset/problem/1633/C)
 - 範例：測資不算太多，直接解即可，須注意生命值最大有 16 位數，然後 go 用 fmt.Scan* 會 TLE，所以全部改用 Scanner。

---

#### [4.Titan - Eating Queries](https://codeforces.com/problemset/problem/1676/E)
 - 範例：題目 time limit per test 給 3.5 秒表示 TLE 機會很高，看到 minimum 可以準備降冪排序，直接解會 TLE，需要用 prefix sum 把排序好的糖質陣列先加好，再用 binary search 去找。

---

#### [5.Odin - 3SUM](https://codeforces.com/problemset/problem/1692/F)
 - 範例：題目只要求三個數加起來後個位數是 3，所以只要算每個個位數的數量，多於三個的只留三個，用剩下的個位數跑三層迴圈找 i j k 就不會 TLE。

---

#### [6.Bahamut - Bit Flipping](https://codeforces.com/problemset/problem/1659/B)
 - 範例：題目要求一定要執行 k 次翻轉，而數字只會有 0 跟 1，可以算出若不選該位數，翻轉 k 次後該位數是 1 還是 0，又相同長度的二進位數字左邊越多 1 越大，所以由左至右讓數字變為 1，若原本數字為 1，k 是奇數，一定要選一次才會維持 1、數字為 0，k 是奇數，不用選最後就會是 1，類推其他兩種狀態，若最後 k 還有剩全部加給最後一位數即可。
