import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import 'bulma/css/bulma.min.css';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import LandingPage from "./pages/LandingPage.tsx";
import DeviceNames from "./pages/DeviceNames.tsx";
import HTTPMethods from "./pages/HTTPMethods.tsx";
import StatusCodes from "./pages/StatusCodes.tsx";
import UserAgents from "./pages/UserAgents.tsx";
import NavbarWrapper from "./components/NavbarWrapper.tsx";




const router = createBrowserRouter([
    {
        path:"/",
        element: <NavbarWrapper/>,
        children:[    {
            path:"/",
            element:<><LandingPage/></>
        },
            {path:"device-names", element:<><DeviceNames/></>},
            {path:"http-methods", element:<><HTTPMethods/></>},
            {path:"status-codes", element:<><StatusCodes/></>},
            {path:"user-agents", element:<><UserAgents/></>},]
    }
])



ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router}/>
  </React.StrictMode>,
)
