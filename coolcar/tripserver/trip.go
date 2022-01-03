package trip

import (
	"context"
	trippd "coolcar/proto/gen/go"
	"fmt"
)

type Service struct  {
	trippd.UnimplementedTripServiceServer
}




func (s Service) GetTrip(c context.Context,req *trippd.GetTripRequest) (*trippd.GetTripResponse,error)  {
	fmt.Println(req.Id)

	response := &trippd.GetTripResponse{
		Id: req.Id,
		Trip: &trippd.Trip{
			Start: "北京",
			StartPos: &trippd.Location{
				Latitude:  101.0,
				Longitude: 101.0,
			},
			PathLocations: nil,
			EndPos: &trippd.Location{
				Latitude:  101.0,
				Longitude: 101.0,
			},
			End:             "天安门",
			DurationSec:     1,
			FeeCent:         1,
			Status:          1,
			IsPromotionTrip: false,
			IsFromGuestUser: false,
		},
	}
	fmt.Println(response)
	return response, nil
}

func main() {

}
