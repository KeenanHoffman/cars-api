package routes

//info := make(chan structs.Car)
//go func() {
//	cars := []structs.Car{
//		{
//			Name: "Charger",
//		},
//		{
//			Name: "Terrain",
//		},
//		{
//			Name: "Camero",
//		},
//	}
//	for _, car := range cars {
//		info <- car
//		time.Sleep(time.Second)
//	}
//}()
//
//
//context.Stream(func(w io.Writer) bool {
//	select {
//	case tm := <-info:
//		context.SSEvent("", tm)
//	}
//	return true
//})
