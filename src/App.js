import Navbar from "./Navbar";
import { BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Home from "./Home";
import Signin from "./Signin";
import Signup from "./Signup";
import Polls from "./Polls";
import { AuthProvider } from './AuthContext';

function App() {
  const isLoggedIn = false; 

  return (
    <AuthProvider>
      <Router>
        <div className="App">
          <Navbar isLoggedIn={isLoggedIn}/>
          <div className="content">
            <Routes>
              <Route exact path="/" element={<Home></Home>}></Route>
              <Route exact path="/signin" element={<Signin></Signin>}></Route>
              <Route exact path="/signup" element={<Signup></Signup>}></Route>
              <Route exact path="/polls" element={<Polls isLoggedIn={isLoggedIn} />}></Route>
            </Routes>
          </div>
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
