#include <bits/stdc++.h>

using namespace std;
typedef long long ll;

int main() {
  vector<int> vec(3);
  for (int i = 0; i < 3; i++) {
    cin >> vec[i];
  }
  sort(vec.begin(), vec.end());
  for (int i = 0; i < 2; i++) {
    cout << vec[i] << " ";
  }
  cout << vec[2] << endl;
}