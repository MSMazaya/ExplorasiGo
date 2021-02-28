import React from 'react'
const Phone = () => {
    return (
    <div className="wadah">
        <form className="form">
          <h1>Gelang Anti Kekerasan</h1>
          <label htmlFor="pnumber">Input your phone number:</label>
          <input type="text" id="pnumber" name="phonenumber" placeholder="Your phone number..."></input>
          <input type="submit" value="Submit"></input>
        </form>
      </div>
    )
}

export default Phone

