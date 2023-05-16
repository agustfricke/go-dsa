function countTo10000() {
  let start = new Date();

  for (let i = 1; i <= 10000000; i++) {
    console.log(i);
  }

  let elapsed = (new Date() - start) / 1000;
  console.log("Tiempo transcurrido: " + elapsed + " segundos");
}

countTo10000();
