<script>
  import {CalculateWorktime} from '../wailsjs/go/main/Time.js'

  let resultText = "...";
  let mustHr = 8;
  let mustMin = 30;
  let hr;
  let min;

  function calculate(){
    CalculateWorktime(hr, min, mustHr, mustMin).then(result => resultText = result)
  }
</script>

<main>
  <!-- <img alt="Wails logo" id="logo" src="{logo}"> -->
  <br />
  <div class="result">How long do you have to work?</div>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={mustHr} class="input" type="number" id="hr"/>Hour
    <input autocomplete="off" bind:value={mustMin} class="input" type="number" id="min"/>Minutes
</div>
<div class="result" id="result"></div>
<p>Please enter your start time 👇 (leave empty when started now)</p>
  <div class="input-box" id="input">
    <input autocomplete="off" bind:value={hr} class="input" type="number" id="hr" placeholder="H"/>Hour
    <input autocomplete="off" bind:value={min} class="input" type="number" id="min" placeholder="M"/>Minutes
</div>

<div class="input-btn">
    <p>{resultText}</p>
    <button class="btn" on:click={calculate}>Show Worktime</button>
</div>
</main>

<style>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  #input{
    display: flex;
    gap: 8px;
    justify-content: center;
    
  }

  .input-btn .btn {
    width: 130px;
    height: 30px;
    border-radius: 3px;
    border: none;
    margin-top: 10px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-btn .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  }

  .input-btn .btn:active {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    transform: scale(0.95);
  }

  .input-box{
    display: flex;
    align-items: center;
  }

  .input-box .input {
    text-align: center;
    border: none;
    border-radius: 3px;
    outline: none;
    height: 22px;
    width: 40px;
    line-height: 30px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
