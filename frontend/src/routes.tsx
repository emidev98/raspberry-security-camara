import {
    HistoryDetails,
    History,
    Login,
    Stream
} from "./pages";

export const routes = () => {
    return [
        {
            path: "/login",
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
            path: "*",
            element: <Login />,
        }
    ]
}
