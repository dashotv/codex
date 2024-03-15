#!/usr/bin/env bash
SOURCE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
set -e

# go run . scraper myanime

# go run . downloader metube http://10.0.4.62:10003/add

# "$SOURCE_DIR/create.sh" "swallowed star" "https://myanime.live/tag/swallowed-star/"
"$SOURCE_DIR/create.sh" "perfect world" "https://myanime.live/tag/perfect-world/" | jq
# "$SOURCE_DIR/create.sh" "a will eternal" "https://myanime.live/tag/a-will-eternal/"
# "$SOURCE_DIR/create.sh" "battle through the heavens" "https://myanime.live/tag/battle-through-the-heavens/"
# "$SOURCE_DIR/create.sh" "honor of kings chapter of glory" "https://myanime.live/tag/honor-of-kings-chapter-of-glory/"
# "$SOURCE_DIR/create.sh" "against the gods" "https://myanime.live/tag/against-the-gods/"
# "$SOURCE_DIR/create.sh" "glorious revenge of ye feng" "https://myanime.live/tag/dubu-wangu/"
# "$SOURCE_DIR/create.sh" "the abyss game" "https://myanime.live/tag/shenyuan-youxi/"
# "$SOURCE_DIR/create.sh" "shrouding the heavens" "https://myanime.live/tag/shrouding-the-heavens/"
# "$SOURCE_DIR/create.sh" "a record of a mortals journey to immortality" "https://myanime.live/tag/fan-ren-xiu-xian-chuan/"
# "$SOURCE_DIR/create.sh" "spare me great lord" "https://myanime.live/tag/spare-me-great-lord/"
# "$SOURCE_DIR/create.sh" "law of devil" "https://myanime.live/tag/law-of-devil/"
# "$SOURCE_DIR/create.sh" "the invincible" "https://myanime.live/tag/shi-fang-wu-sheng/"
# "$SOURCE_DIR/create.sh" "the proud emperor of eternity" "https://myanime.live/tag/the-proud-emperor-of-eternity/"
# "$SOURCE_DIR/create.sh" "legend of xianwu" "https://myanime.live/tag/legend-of-xianwu/"
# "$SOURCE_DIR/create.sh" "martial god asura" "https://myanime.live/tag/martial-god-asura/"
# "$SOURCE_DIR/create.sh" "100000 years of refining qi" "https://myanime.live/tag/100-000-years-of-refining-qi/page/10/"
# "$SOURCE_DIR/create.sh" "the eternal strife" "https://myanime.live/tag/bu-shi-bu-mie/"
# "$SOURCE_DIR/create.sh" "the great ruler" "https://myanime.live/tag/the-great-ruler/"

"$SOURCE_DIR/api.sh" "PageService.Index" '{}' | jq
# "$SOURCE_DIR/find.sh" "aadb2fbf0021166563f0a72c16248c33"
