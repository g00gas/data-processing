import Navbar from "./Navbar.tsx";
import {Outlet} from "react-router-dom";

const NavbarWrapper = () => {
    return(<div>
        <Navbar/>
        <Outlet/>
    </div>)
}

export default NavbarWrapper