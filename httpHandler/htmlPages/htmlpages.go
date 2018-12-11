package htmlPages

const Formpage = `
  <h1>Login</h1>
  <form method="POST" action="/login">
      <label for="name">name</label>
      <input type="text" id="name" name="name">
      <label for="id">id</label>
      <input type="text" id="id" name="id">
	  <label for="age">age</label>
      <input type="text" id="age" name="age">
      <button type="submit">Login</button>
      <button type="submit">register</button>
      <button type="submit">fetch post</button>

  </form>
 `

const Registerpage = `
  <h1>register</h1>
  <form method="GET" action="/">
      <label for="name">name</label>
      <input type="text" id="name" name="name">
      <label for="id">id</label>
      <input type="text" id="id" name="id">
	  <label for="age">age</label>
      <input type="text" id="age" name="age">
      <button type="submit">register</button>
  </form>
 `

const InternalPage = `
 <h1>Internal</h1>
 <hr>
 <small>User: %s</small>
 <form method="GET" action="/logout">
     <label for="title">title</label>
     <input type="text" id="title" name="title">
     <label for="discription">discription</label>
     <input type="text" id="discription" name="discription">
      <button type="submit">post</button>

      <button type="submit">Logout</button>
 </form>
 `

const Fetchform = `
<h1>fetch details<h1>
<form method="GET">
        <label for="name">name</label>
        <input type="text" id="name" name="name">
        <button type="submit">fetch</button>
</form>
	`
