const express = require('express');
const sequelize = require('./db');
const { Sequelize } = require('sequelize');
const User = require('./models/user');
const morgan = require('morgan');
const app = express();
const PORT = 3000;


app.use((req, res, next) => {
    const start = process.hrtime();

    res.on('finish', () => {
        const [seconds, nanoseconds] = process.hrtime(start);
        const elapsed = (seconds * 1000 + nanoseconds / 1e6).toFixed(3); 
        console.log(`Request to ${req.method} ${req.url} took ${elapsed} ms`);
    });

    next();
});

app.use(morgan('combined'));

app.use(express.json());

// Синхронизация базы данных
sequelize.sync()
  .then(() => console.log('Database synced'))
  .catch((err) => console.error('Error syncing database:', err));

function heavyMemoryOperation() {
    const size = 10000000;
    const data = new Array(size);

    // Fill the array with random values
    for (let i = 0; i < size; i++) {
        data[i] = i * 2
    }

    // Perform some operation (e.g., calculate the sum)
    let sum = 0;
    for (let i = 0; i < size; i++) {
        sum += data[i];
    }

    // Release memory by clearing the array
    data.length = 0;

    return sum;
}

app.get('/health', (req, res) => {
    res.json({message: "OK"})
})

// Route that triggers the heavyMemoryOperation
app.post('/upload', (req, res) => {
    const result = heavyMemoryOperation();
    res.json({ result });
});

app.post('/users', async (req, res) => {
  try {
      const { password, name, role, email } = req.body;
      // SQL-запрос на создание пользователя, идентичный Go/GORM
      const [user] = await sequelize.query(
          `INSERT INTO "users" ("name", "email", "role", "password") 
           VALUES ($1, $2, $3, $4) RETURNING "id"`,
          {
              bind: [name, email, role, password], // Параметры запроса
              type: Sequelize.QueryTypes.INSERT, // Тип запроса
          }
      );
      res.status(201).json(user[0]);
  } catch (error) {
      res.status(500).json({ error: `Failed to create user: ${error.message}` });
  }
});

app.get('/users', async (req, res) => {
  try {
      // SQL-запрос на получение списка пользователей, идентичный Go/GORM
      const [users] = await sequelize.query(
          `SELECT * FROM "users" LIMIT $1`,
          {
              bind: [20], // Лимит выборки, идентичный Go/GORM
              type: Sequelize.QueryTypes.SELECT, // Тип запроса
          }
      );
      res.status(200).json(users);
  } catch (error) {
      res.status(500).json({ error: `Failed to retrieve users: ${error.message}` });
  }
});


app.delete('/users/:id', async (req, res) => {
    try {
      const { id } = req.params;
      const deleted = await User.destroy({ where: { id } });
      if (deleted) {
        res.status(204).send();
      } else {
        res.status(404).json({ error: 'User not found' });
      }
    } catch (error) {
      res.status(500).json({ error: 'Failed to delete user' });
    }
});
  
  
  

app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
