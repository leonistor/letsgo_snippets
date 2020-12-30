const { openBrowser, goto, click, textBox, focus, write, button, closeBrowser } = require('taiko');
(async () => {
  try {
    await openBrowser();
    await goto("https://localhost:8080/");
    await click("Login");
    await focus(textBox("email"));
    await write("zuza@test.com");
    await focus(textBox("password"));
    await write("test");
    await focus(button("Login"));
    await click(button("Login"));
    await click("Home");
    await click("Logout");
  } catch (error) {
    console.error(error);
  } finally {
    await closeBrowser();
  }
})();
