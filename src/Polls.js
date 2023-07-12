import React, { useEffect, useState, useContext } from 'react';
import { useLocation, useHistory } from 'react-router-dom';
import { AuthContext } from './AuthContext';

const Polls = () => {
  const location = useLocation();
  // const userData = location.state[0];

  const storedUserData = localStorage.getItem('userData');
  const userData = storedUserData ? JSON.parse(storedUserData)[0] : null;

  const [polls, setPolls] = useState([]);
  const [selectedOption, setSelectedOption] = useState(null);
  const { setIsLoggedIn } = useContext(AuthContext);
  const [voteCounts, setVoteCounts] = useState({});
  const [newPollQuestion, setNewPollQuestion] = useState('');
  const [newPollOptions, setNewPollOptions] = useState([]);
  const [creatingPoll, setCreatingPoll] = useState(false);

  useEffect(() => {
    setIsLoggedIn(true);
    fetch('http://localhost:8080/api/v1/polls')
      .then((response) => response.json())
      .then((data) => {
        setPolls(data);
      })
      .catch((error) => console.error('Error fetching polls:', error));
  }, []);

  const fetchVoteCounts = (poll) => {
    const pollOptionIds = poll.edges.poll_options.map((option) => option.id);
    Promise.all(
      pollOptionIds.map((pollOptionId) =>
        fetch(`http://localhost:8080/api/v1/votes/${pollOptionId}`)
          .then((response) => response.json())
      )
    )
      .then((data) => {
        const counts = pollOptionIds.reduce((acc, pollOptionId, index) => {
          const pollOptionVotes = data[index];
          const count = pollOptionVotes.length;
          const userIds = pollOptionVotes.map((vote) => vote.userId);
          acc[pollOptionId] = { count, userIds };
          return acc;
        }, {});

        setVoteCounts((prevVoteCounts) => ({
          ...prevVoteCounts,
          [poll.id]: counts,
        }));
      })
      .catch((error) => console.error(`Error fetching votes for poll ${poll.id}:`, error));
  };

  const fetchPolls = () => {
    fetch('http://localhost:8080/api/v1/polls')
      .then((response) => response.json())
      .then((data) => {
        setPolls(data);
      })
      .catch((error) => console.error('Error fetching polls:', error));
  };

  const handleCreatePoll = () => {
    setCreatingPoll(true);
    const newPollData = {
      question: newPollQuestion,
      options: newPollOptions,
    };
    fetch('http://localhost:8080/api/v1/polls', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newPollData),
    })
      .then((response) => {
        if (response.ok) {
          console.log('Poll created');
          setNewPollQuestion('');
          setNewPollOptions([]);
          setCreatingPoll(false);
          fetchPolls();
        } else {
          console.error('Error creating poll');
          setCreatingPoll(false);
        }
      })
      .catch((error) => {
        console.error('Error creating poll:', error);
        setCreatingPoll(false);
      });
  };

  const handleVoteSubmit = (event, poll) => {
    event.preventDefault();
    const voteData = {
      userID: userData.id.toString(),
      pollOptionID: selectedOption,
    };

    console.log(JSON.stringify(voteData));
    fetch('http://localhost:8080/api/v1/votes', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(voteData),
    })
      .then((response) => {
        if (response.ok) {
          // Vote successfully submitted
          console.log('Vote submitted');
          fetchVoteCounts(poll); // Fetch updated vote counts
        } else {
          // Handle the error case
          console.error('Error submitting vote');
        }
      })
      .catch((error) => console.error('Error submitting vote:', error));
  };

  const handleOptionChange = (event) => {
    setSelectedOption(event.target.value);
  };

  const handleVoteCountClick = (pollId, optionId) => {
    const userIds = voteCounts[pollId][optionId].userIds;
    // Fetch user data using user IDs
    Promise.all(
      userIds.map((userId) =>
        fetch(`http://localhost:8080/api/v1/users/${userId}`)
          .then((response) => response.json())
      )
    )
      .then((userData) => {
        const userFirstNames = userData.map((user) => user[0].firstName);
        console.log('User First Names:', userFirstNames);

        // Render the user first names beside the vote count
        setVoteCounts((prevVoteCounts) => ({
          ...prevVoteCounts,
          [pollId]: {
            ...prevVoteCounts[pollId],
            [optionId]: {
              ...prevVoteCounts[pollId][optionId],
              userFirstNames: userFirstNames,
            },
          },
        }));
      })
      .catch((error) =>
        console.error(`Error fetching user data for user IDs: ${userIds}`, error)
      );
  };

  useEffect(() => {
    console.log(selectedOption);
  }, [selectedOption]);

  return (
    <div>
      {/* Render the user information */}
      <h2>Welcome, {userData.firstName}!</h2>
      <br />
      <h3>Create New Poll</h3>
      <form className='create'>
        <div>
          <label htmlFor="newPollQuestion">Question:</label>
          <input
            type="text"
            id="newPollQuestion"
            value={newPollQuestion}
            onChange={(e) => setNewPollQuestion(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="newPollOptions">Options:</label>
          <input
            type="text"
            id="newPollOptions"
            value={newPollOptions.join(',')}
            onChange={(e) => setNewPollOptions(e.target.value.split(','))}
          />
        </div>
        <button
          type="button"
          onClick={handleCreatePoll}
          disabled={creatingPoll || !newPollQuestion || newPollOptions.length < 2}
        >
          {creatingPoll ? 'Creating Poll...' : 'Create Poll'}
        </button>
      </form>
      <br />
      {/* Render the polls and options */}
      {polls && polls.length > 0 ? (
        polls.map((poll) => (
          <div key={poll.id}>
            <h3>{poll.question}</h3>
            {poll.edges && poll.edges.poll_options && poll.edges.poll_options.length > 0 ? (
              <form onSubmit={(event) => handleVoteSubmit(event, poll)}>
                {poll.edges.poll_options.map((option) => (
                  <label className="radio" key={option.id}>
                  <input
                    type="radio"
                    name={`poll_${poll.id}`}
                    value={option.id}
                    checked={selectedOption === option.id}
                    onChange={handleOptionChange}
                  />
                  {option.option}
                  {selectedOption === option.id && <span className="dot"></span>}
                  {voteCounts[poll.id] && voteCounts[poll.id][option.id] && (
                    <span
                      className="vote-count"
                      onClick={() => handleVoteCountClick(poll.id, option.id)}
                    >
                      ({voteCounts[poll.id][option.id].count} vote(s))
                    </span>
                  )}
                  {voteCounts[poll.id] &&
                    voteCounts[poll.id][option.id] &&
                    voteCounts[poll.id][option.id].userFirstNames &&
                    voteCounts[poll.id][option.id].userFirstNames.length > 0 && (
                      <span className="user-names">
                        - {voteCounts[poll.id][option.id].userFirstNames.join(', ')}
                      </span>
                    )}
                </label>
                ))}
                <button type="submit" disabled={selectedOption === null}>
                  Submit Vote
                </button>
              </form>
            ) : (
              <p>No options available for this poll.</p>
            )}
            <br />
          </div>
        ))
      ) : (
        <p>No polls available.</p>
      )}
    </div>
  );
};

export default Polls;
