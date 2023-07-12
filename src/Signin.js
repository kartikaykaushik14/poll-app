import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

const Signin = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isPending, setIsPending] = useState(false);

    const navigate = useNavigate();

    const handleSignIn = (e) => {
        e.preventDefault();
        const userInfo = {username, password};
        setIsPending(true);
        fetch(`http://localhost:8080/api/v1/user/${username}/${password}`)
        .then((response) => {
            setIsPending(false);
            if (response.ok) {
                return response.json();
            } else {
                // Handle authentication failure
                // For example, show an error message
                console.log("Authentication failed");
                alert("Invalid username or password");
                // Clear the input fields
                setUsername("");
                setPassword("");
                throw new Error("Authentication failed");
            }
        })
        .then((data) => {
            if (data.length === 1) {
                // navigate(`/polls`, { state: data });
                localStorage.setItem('userData', JSON.stringify(data));
                navigate(`/polls`);
            } else {
                // Handle authentication failure
                // For example, show an error message
                console.log("Authentication failed");
                alert("Invalid username or password");
                // Clear the input fields
                setUsername("");
                setPassword("");
                throw new Error("Authentication failed");
            }
        })
        .catch((error) => {
            console.error("Error occurred during authentication:", error);
            // Handle error case
            // For example, show an error message
            alert("An error occurred during authentication. Please try again later.");
        });
    }

    return (
        <div className="signin">
          <h2>Existing User? Sign In</h2>
          <form onSubmit={handleSignIn}>
              <label>Username: </label>
              <input
                  type="text"
                  required
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}>
              </input>
              <label>Password: </label>
              <input
                  type="text"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}>
              </input>
              {!isPending && <button>Sign In</button>}
              {isPending && <button disabled>Signing In...</button>}
            </form>

            <h2>New User? <Link to = "/signup">Sign Up</Link></h2>
        </div>
      );
}
 
export default Signin;