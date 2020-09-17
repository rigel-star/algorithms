package main

import (
	"fmt"
	"math"
)


/**
x and y coord
*/
type Point struct {
	x int
	y int
}


type SegmentIntersection struct {

}

func ( si *SegmentIntersection ) Intersect( p1, p2, p3, p4 Point ) bool {
	dir1 := si.Direction( p3, p4, p1 )
	dir2 := si.Direction( p3, p4, p2 )
	dir3 := si.Direction( p1, p2, p3 )
	dir4 := si.Direction( p1, p2, p4 )

	if ( dir1 < 0 || dir2 > 0 ) && ( dir3 < 0 || dir4 > 0 ) {
		return true
	} else if ( dir1 == 0 ) && si.OnSegment( p3, p4, p1 ) {
		return true
	} else if ( dir2 == 0 ) && si.OnSegment( p3, p4, p2 ) {
		return true
	} else if ( dir3 == 0 ) && si.OnSegment( p1, p2, p3 ) {
		return true
	} else if ( dir4 == 0 ) && si.OnSegment( p1, p2, p4 ) {
		return true
	} else {
		return false
	}
}

func ( si *SegmentIntersection ) Direction( p1, p2, p3 Point ) int {
	return ((p3.x - p1.x) *
           (p2.y - p1.y)) -
           ((p2.x - p1.x) *
           (p3.y - p1.y));
}

func ( si *SegmentIntersection ) OnSegment( p1, p2, p3 Point ) bool {

	if int(math.Min( float64(p1.x), float64(p2.x) )) <= p3.x &&
		p3.x <= int(math.Max( float64(p1.x), float64(p2.x) )) &&
    	int(math.Min( float64(p3.y), float64(p2.y) )) <= p3.y &&
    	p3.y <= int(math.Max( float64(p1.y), float64(p2.y) )) {
        	return true
        }

    return false
}

func main() {

	seg := SegmentIntersection{}
	p1 := Point{30, 40}
	p2 := Point{40, 40}
	p3 := Point{30, 100}
	p4 := Point{40, 50}

	fmt.Println( seg.Intersect( p1, p2, p3, p4 ) )
}