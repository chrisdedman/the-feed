/*
 * This Function is used to toggle the display of the feed items.
  * When the user clicks on a feed title, the feed item will be displayed.
  * When the user clicks on the feed title again, the feed item will be hidden.
  * This function is called when the page is loaded.
*/
function toggleFeedItem() {
  var feedTitles = document.querySelectorAll('.clickable');

  feedTitles.forEach(function(feedTitle) {
    feedTitle.addEventListener('click', function() {
      // Toggle the display of the next sibling feed item
      var feedItem = this.nextElementSibling;

      while (feedItem && feedItem.classList.contains('feed-item')) {
        if (feedItem.style.display === 'none' || feedItem.style.display === '') {
          feedItem.style.display = 'block';
        } else {
          feedItem.style.display = 'none';
        }
        feedItem = feedItem.nextElementSibling;
      }
    });
  });
}

/* When the user clicks on the button,
toggle between hiding and showing the dropdown content */
function settings() {
  document.getElementById("setting").classList.toggle("show");
}

// Close the dropdown menu if the user clicks outside of it
window.onclick = function(event) {
  if (!event.target.matches('.dropbtn')) {
    var dropdowns = document.getElementsByClassName("dropdown-content");
    for (var i = 0; i < dropdowns.length; i++) {
      var openDropdown = dropdowns[i];
      if (openDropdown.classList.contains('show')) {
        openDropdown.classList.remove('show');
      }
    }
  }
}

/* 
  setmenu() function is used to set menu items dynamically
  the user to the selected menu option from the frontend.
*/
function setMenu(menuItems) {
  const menuContainer = document.getElementById('menu');
  menuContainer.innerHTML = ''; // Clear existing menu items

  // Loop through menuItems array and create buttons for each item
  menuItems.forEach(item => {
    const button = document.createElement('button');
    button.classList.add('menu');
    button.textContent = item.text;
    button.onclick = () => {
      window.location.href = item.url;
    };
    menuContainer.appendChild(button);
  });
}

/*
  togglePasswordFields() function is used to show/hide the confirm password field
  when the user checks the "New Password" checkbox.
*/
function togglePasswordFields() {
  var passwordDiv = document.getElementById("confirmPasswordDiv");
  var checkbox = document.getElementById("newPassword");
  passwordDiv.style.display = checkbox.checked ? "block" : "none";
}

/*
  passwordRequirements() function is used to check if the password meets the requirements
  and matches the confirm password. If the password does not meet the requirements,
  an error message will be displayed.
*/  
function passwordRequirements(password, confirmPassword) {
  // Check if password meets requirements and matches confirm password
  if (password.length < 8 || !/\d/.test(password) || !/[A-Z]/.test(password) || !/[a-z]/.test(password) || password !== confirmPassword) {
    document.getElementById('message').textContent = 'Password does not meet requirements.';
    return false;
  }
  return true;
}

/*
  Validate password. Check for at least 8 characters, 
  one number, one uppercase letter, one lowercase letter, 
  and that the password matches the confirm password.
*/
function validatePassword() {
  const password           = document.getElementById('password').value;
  const confirmPassword    = document.getElementById('confirmPassword').value;
  const charLengthCheck    = document.getElementById('charLengthCheck');
  const numberCheck        = document.getElementById('numberCheck');
  const upperCaseCheck     = document.getElementById('upperCaseCheck');
  const lowerCaseCheck     = document.getElementById('lowerCaseCheck');
  const passwordMatchCheck = document.getElementById('passwordMatchCheck');

  // Check password length
  if (password.length >= 8) {
    charLengthCheck.textContent = '✓';
  } else {
    charLengthCheck.textContent = '✗';
  }

  // Check for at least one number
  if (/\d/.test(password)) {
    numberCheck.textContent = '✓';
  } else {
    numberCheck.textContent = '✗';
  }

  // Check for at least one uppercase letter
  if (/[A-Z]/.test(password)) {
    upperCaseCheck.textContent = '✓';
  } else {
    upperCaseCheck.textContent = '✗';
  }

  // Check for at least one lowercase letter
  if (/[a-z]/.test(password)) {
    lowerCaseCheck.textContent = '✓';
  } else {
    lowerCaseCheck.textContent = '✗';
  }

  // Check if password matches confirm password
  if (password === confirmPassword) {
    passwordMatchCheck.textContent = '✓';
  } else {
    passwordMatchCheck.textContent = '✗';
  }
}