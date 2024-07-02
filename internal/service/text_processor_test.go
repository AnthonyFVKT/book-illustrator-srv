package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSplitText(t *testing.T) {
	processor := NewProcessor(3)

	// Test case 1: Sentences with different punctuation marks
	text := `Hello, world! How are you? I'm fine, thank you. Have a nice day.`
	expectedGroups := []string{"Hello, world. How are you. I'm fine, thank you.", "Have a nice day."}
	require.Equal(t, expectedGroups, processor.SplitText(text))

	// Test case 2: Sentences with multiple punctuation marks
	text = "This is a test. This is another test. And this is a third test. Let's go."
	expectedGroups = []string{"This is a test. This is another test. And this is a third test.", "Let's go."}
	require.Equal(t, expectedGroups, processor.SplitText(text))

	// Test case 3: Sentences with no punctuation marks
	text = "This is a sentence. This is another sentence. This is a third sentence."
	expectedGroups = []string{"This is a sentence. This is another sentence. This is a third sentence."}
	require.Equal(t, expectedGroups, processor.SplitText(text))

	// Test case 4: Sentences with a single punctuation mark
	text = "This is a sentence. This is another sentence."
	expectedGroups = []string{"This is a sentence. This is another sentence."}
	require.Equal(t, expectedGroups, processor.SplitText(text))

	// Test case 5: Sentences with a mixture of punctuation marks and different lengths
	text = "This is a short sentence. This is a longer sentence. This is a very long sentence."
	expectedGroups = []string{"This is a short sentence. This is a longer sentence. This is a very long sentence."}
	require.Equal(t, expectedGroups, processor.SplitText(text))
}
