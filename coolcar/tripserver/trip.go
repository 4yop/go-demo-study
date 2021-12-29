package trip

import (
	"context"
	trippd "coolcar/proto/gen/go"
)

type Service struct  {
	trippd.UnimplementedTripServiceServer
}




func (s Service) GetTrip(c context.Context,req *trippd.GetTripRequest) (*trippd.GetTripResponse,error)  {
	return &trippd.GetTripResponse{
		Id: req.Id,
		Trip: &trippd.Trip{
			Start: "",
			StartPos: &trippd.Location{
				Latitude:  0,
				Longitude: 0,
			},
			PathLocations: nil,
			EndPos: &trippd.Location{
				Latitude:  0,
				Longitude: 0,
			},
			End:             "",
			DurationSec:     0,
			FeeCent:         0,
			Status:          0,
			IsPromotionTrip: false,
			IsFromGuestUser: false,
		},
	}, nil
}

func main() {

}
