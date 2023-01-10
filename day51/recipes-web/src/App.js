import React from 'react';
import logo from './logo.svg';
import './App.css';
import './Recipe.css';

const Recipe = (props) => (
  <div className='recipe'>
    <h4>{props.recipe.name}</h4>
    <ul>
      {props.recipe.ingredients && props.recipe.ingredients.map((ing, ind) => (
        <li>{ing}</li>
      ))}
    </ul>
  </div>
)

function App() {
  let recipes = [
    {
      "name": "Oregano Marinated Chicken",
      "tags": [
        "main",
        "chicken"
      ],
      "ingredients": [],
      "instructions": []
    },
    {
      "name": "Green pea soup with cheddar scallion panini",
      "tags": [
        "soup",
        "main",
        "panini"
      ],
      "ingredients": [],
      "instructions": []
    }
  ] 

  return (
    <div>
      {recipes.map((recipe, index) => (
        <Recipe recipe={recipe} />
      ))}
    </div>
  );
}

export default App;
