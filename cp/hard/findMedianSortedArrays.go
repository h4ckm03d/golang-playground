package hard

import (
	"math"
)

/*



 double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
    int N1 = nums1.size();
    int N2 = nums2.size();
    if (N1 < N2) return findMedianSortedArrays(nums2, nums1);	// Make sure A2 is the shorter one.

    int lo = 0, hi = N2 * 2;
    while (lo <= hi) {
        int mid2 = (lo + hi) / 2;   // Try Cut 2
        int mid1 = N1 + N2 - mid2;  // Calculate Cut 1 accordingly

        double L1 = (mid1 == 0) ? INT_MIN : nums1[(mid1-1)/2];	// Get L1, R1, L2, R2 respectively
        double L2 = (mid2 == 0) ? INT_MIN : nums2[(mid2-1)/2];
        double R1 = (mid1 == N1 * 2) ? INT_MAX : nums1[(mid1)/2];
        double R2 = (mid2 == N2 * 2) ? INT_MAX : nums2[(mid2)/2];

        if (L1 > R2) lo = mid2 + 1;		// A1's lower half is too big; need to move C1 left (C2 right)
        else if (L2 > R1) hi = mid2 - 1;	// A2's lower half too big; need to move C2 left.
        else return (max(L1,L2) + min(R1, R2)) / 2;	// Otherwise, that's the right cut.
    }
    return -1;
}
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	if n1 < n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	lo, hi := 0, n2*2
	for lo <= hi {
		mid2 := (lo + hi) / 2
		mid1 := n1 + n2 - mid2
		l1 := math.MinInt64
		if mid1 != 0 {
			l1 = nums1[(mid1-1)/2]
		}
		l2 := math.MinInt64
		if mid2 != 0 {
			l2 = nums2[(mid2-1)/2]
		}
		r1 := math.MaxInt64
		if mid1 != n1*2 {
			r1 = nums1[mid1/2]
		}
		r2 := math.MaxInt64
		if mid2 != n2*2 {
			r2 = nums2[mid2/2]
		}
		if l1 > r2 {
			lo = mid2 + 1
		} else if l2 > r1 {
			hi = mid2 - 1
		} else {
			return float64(max(l1, l2)+min(r1, r2)) / 2.0
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
