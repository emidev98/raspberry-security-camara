import {
    HistoryDetails,
    History,
    Login,
    Stream,
    Error
} from "./pages";

export const routes = () => {
    return [
        {
            path: "/",
            element: <Login />,
        },
        {
            path: "/stream",
            element: <Stream />,
        },
        {
            path: "/history",
            element: <History />,
        },
        {
            path: "/history/{id}",
            element: <HistoryDetails />,
        },
        {
            path: "/error",
            element: <Error msg="API Error"/>,
        },
        // {
        //     path: "*",
        //     element: <Error msg="Page not found"/>,
        // },
    ]
}
