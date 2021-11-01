![GitHub license](https://img.shields.io/github/license/torresflo/steam-box-for-readme.svg)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
![GitHub contributors](https://img.shields.io/github/contributors/torresflo/steam-box-for-readme.svg)
![GitHub issues](https://img.shields.io/github/issues/torresflo/steam-box-for-readme.svg)

<p align="center">
  <h1 align="center">Steam Box for readme</h3>

  <p align="center">
    A little fork of <a href="https://github.com/YouEclipse/steam-box">steam-box</a> to help you update your readme file with Steam playtime leaderboard and Steam recent played games.
    <br />
    <a href="https://github.com/torresflo/steam-box-for-readme/issues">Report a bug or request a feature</a>
  </p>
</p>

## Table of Contents

* [Installation](#installation)
   * [Preparation work](#preparation-work)
   * [Project setup](#project-setup)
* [Example](#example)
* [Contributing](#contributing)
* [License](#license)

---

## Installation

### Preparation work
1. Create a Steam  API key. (https://steamcommunity.com/dev/apikey)
1. Find the steam ID (steamID64) of your account. (https://steamid.io)
1. Add comments to the place where you want to update in your markdown file.

You can add your most played games leaderboard with:
   ```markdown
    <!-- steam-box-playtime start -->
    <!-- steam-box-playtime end -->
   ```

And you can add your recent played games leaderboard with:
   ```markdown
    <!-- steam-box-recent start -->
    <!-- steam-box-recent end -->
   ```

### Project setup

For updating your github profile README, you can follow [steam-box.yml](https://github.com/torresflo/steam-box-for-readme/blob/master/.github/workflows/schedule.yml) to create a GitHub Action in your README repository.

1. Go to the repo **Settings > Secrets**
1. Add the following environment variables:
   - **STEAM_API_KEY:** The steam API key you created above. 
   - **STEAM_ID:** The steam ID of your account.
1. If you want to show specific games,put the ids in environmet variable **APP_ID**:
   - like `APP_ID=431960,730`
   - you can get the id of a game from the store url: `https://store.steampowered.com/app/`**730**`/CounterStrike_Global_Offensive/`

## Example

Here is an example from my own profile that you could obtain:

![Example image](https://raw.githubusercontent.com/torresflo/steam-box-for-readme/master/images/1.png)

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License
Distributed under the [Apache-2.0](./LICENSE) License. See `LICENSE` for more information.
