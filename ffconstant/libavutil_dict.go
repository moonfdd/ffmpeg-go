package ffconstant

/**
  * @addtogroup lavu_dict AVDictionary
  * @ingroup lavu_data
  *
  * @brief Simple key:value store
  *
  * @{
  * Dictionaries are used for storing key:value pairs. To create
  * an AVDictionary, simply pass an address of a NULL pointer to
  * av_dict_set(). NULL can be used as an empty dictionary wherever
  * a pointer to an AVDictionary is required.
  * Use av_dict_get() to retrieve an entry or iterate over all
  * entries and finally av_dict_free() to free the dictionary
  * and all its contents.
  *
  @code
    AVDictionary *d = NULL;           // "create" an empty dictionary
    AVDictionaryEntry *t = NULL;

    av_dict_set(&d, "foo", "bar", 0); // add an entry

    char *k = av_strdup("key");       // if your strings are already allocated,
    char *v = av_strdup("value");     // you can avoid copying them like this
    av_dict_set(&d, k, v, AV_DICT_DONT_STRDUP_KEY | AV_DICT_DONT_STRDUP_VAL);

    while (t = av_dict_get(d, "", t, AV_DICT_IGNORE_SUFFIX)) {
        <....>                             // iterate over all entries in d
    }
    av_dict_free(&d);
  @endcode
*/

const AV_DICT_MATCH_CASE = 1    /**< Only get an entry with exact-case key match. Only relevant in av_dict_get(). */
const AV_DICT_IGNORE_SUFFIX = 2 /**< Return first entry in a dictionary whose first part corresponds to the search key,
  ignoring the suffix of the found key string. Only relevant in av_dict_get(). */
const AV_DICT_DONT_STRDUP_KEY = 4 /**< Take ownership of a key that's been
  allocated with av_malloc() or another memory allocation function. */
const AV_DICT_DONT_STRDUP_VAL = 8 /**< Take ownership of a value that's been
  allocated with av_malloc() or another memory allocation function. */
const AV_DICT_DONT_OVERWRITE = 16 ///< Don't overwrite existing entries.
const AV_DICT_APPEND = 32         /**< If the entry already exists, append to it.  Note that no
  delimiter is added, the strings are simply concatenated. */
const AV_DICT_MULTIKEY = 64 /**< Allow to store several equal keys in the dictionary */
