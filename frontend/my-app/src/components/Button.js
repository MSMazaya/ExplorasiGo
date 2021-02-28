import React from 'react'
import axios from 'axios'
const Button = () => {

    const apiCall = () => {
        axios.get('http://localhost:8080/get/6')
    }
    return (
    <div className="jumbotron">
        <h1>Tombol Darurat</h1>
        <br></br>
        <p>Tombol Merah akan mengaktifkan suara sirine dan mengirimkan pesan ke kontak yang sudah didaftarkan. Tombol Hijau akan mengaktifkan perekaman suara dan mengirimkan pesan ke kontak yang sudah didaftarkan.</p>
        <div class="buttons">
            <div onClick={()=>{apiCall()}} class="Red">Merah</div>
            <div onClick={()=>{apiCall()}} class="Green">Hijau</div>
        </div>
    </div>
    )
}

export default Button
