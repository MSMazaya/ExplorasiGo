import './components/style2.css'
import React,{useEffect} from 'react'

function App() {
  
  const accountSid = "AC45d9bf7b4edeb91dbbd5ae152be0fe58";
  const authToken = "a2d456698b19b0e14a564a1d593e4ad5";
  const client = require('twilio')(accountSid, authToken);

  useEffect(() => { 
      client.messages
      .create({
        body: 'This is the ship that made the Kessel Run in fourteen parsecs?',
        from: '+16152819135',
        to: '+628112121326'
      })
    .then(message => console.log(message.sid))
  }, [])

  return (
      <form class="form">
          <h1>Gelang Anti Kekerasan</h1>
          <label htmlFor="pnumber">Input your phone number:</label>
          <input type="text" id="pnumber" name="phonenumber" placeholder="Masukkan Nomor Handphone"></input>
          <input type="submit" onClick={()=>{}}></input>
      </form>
  );
}

export default App;
