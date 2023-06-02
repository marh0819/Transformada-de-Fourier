
fetch('http://localhost:8080/transformada')
  .then(response => response.json())
  .then(data => {    // Obtener los datos necesarios para la gráfica (ejemplo: etiquetas y valores)
   /* const labels = data.map(item => item.label);
    const values = data.map(item => item.value);

    // Crear la configuración de la gráfica
    const chartConfig = {
      type: 'line',
      data: {
        labels: labels,
        datasets: [{
          label: 'Valores',
          data: values,
          backgroundColor: 'rgba(0, 123, 255, 0.5)', // Color de fondo de las barras
          borderColor: 'rgba(0, 123, 255, 1)', // Color del borde de las barras
          borderWidth: 1 // Ancho del borde de las barras
        }]
      },
      options: {
        responsive: true,
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    };*/
      
      var fase = data.map(function (item) {
        return item.fase;
      });
      
      var mag = data.map(function (item) {
        return item.mag;
      });
      var ctx = document.getElementById('myChart').getContext('2d');
      
      var chart = new Chart(ctx, {
        type: 'line',
        data: {
          labels: fase, // Fase como eje X
          datasets: [{
            label: 'Magnitud',
            data: mag, // Magnitud como eje Y
            backgroundColor: 'rgba(41,10,131,0.5)',
            borderColor: 'rgba(41,10,131,0.5)',
            borderWidth: 4
          }/*,
          {
            label: 'fases',
            data: fase, // Magnitud como eje Y
            backgroundColor: 'rgba(0, 123, 255, 0.5)',
            borderColor: 'rgba(0, 123, 255, 1)',
            borderWidth: 1
          }*/]
        },
        options: {
          responsive: true,
          scales: {
            x: {
              display: true,
              title: {
                display: true,
                text: 'Fase'
              }
            },
            y: {
              display: true,
              title: {
                display: true,
                text: 'Magnitud'
              }
            }
          }
        }
      });

    
  })
  .catch(error => {
    console.log('Error al obtener los datos:', error);
  });