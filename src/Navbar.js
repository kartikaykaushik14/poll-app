import { Link } from "react-router-dom";
import { AuthContext } from './AuthContext';
import React, { useContext } from 'react';

const Navbar = () => {
    const { isLoggedIn } = useContext(AuthContext);

    return ( 
        <nav className="navbar">
            <h1>The Poll App</h1>
            <div className="links">
                <Link to = "/"> Home </Link>
                {isLoggedIn ? (<Link to="/signin">Sign Out</Link>) : (<Link to="/signin">Sign In</Link>)}
            </div>
        </nav>
     );
}
 
export default Navbar;