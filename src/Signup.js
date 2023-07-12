import { useState } from "react";
import { useNavigate } from "react-router-dom";

const Signup = () => {
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isPending, setIsPending] = useState(false);
    
    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();
        const user = {firstName, lastName, username, password};
        console.log(user);
        setIsPending(true);
        fetch('http://localhost:8080/api/v1/user', {
            method: 'POST',
            headers: {"Content-Type":"application/json"},
            body:JSON.stringify(user)
        }).then(() => {
            setIsPending(false);
            navigate(`/signin`);
        })
    }

    return (
        <div className="create">
          <h2>Enter your details</h2>
          <form onSubmit={handleSubmit}>
              <label>First Name: </label>
              <input
                  type="text"
                  required
                  value={firstName}
                  onChange={(e) => setFirstName(e.target.value)}>
              </input>
              <label>Last Name: </label>
              <input
                  type="text"
                  required
                  value={lastName}
                  onChange={(e) => setLastName(e.target.value)}>
              </input>
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
              {!isPending && <button>Create User</button>}
              {isPending && <button disabled>Creating User...</button>}
          </form>
        </div>
      );
}
 
export default Signup;