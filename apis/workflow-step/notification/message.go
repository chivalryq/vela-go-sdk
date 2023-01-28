/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"vela-go-sdk/apis/utils"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Message{}

// Message Specify the message that you want to sent, refer to [dingtalk messaging](https://developers.dingtalk.com/document/robots/custom-robot-access/title-72m-8ag-pqw)
type Message struct {
	actionCard NullableActionCard `json:"actionCard,omitempty"`
	at         NullableAt         `json:"at,omitempty"`
	feedCard   *FeedCard          `json:"feedCard,omitempty"`
	link       *Link              `json:"link,omitempty"`
	markdown   NullableMarkdown   `json:"markdown,omitempty"`
	// msgType can be text, link, mardown, actionCard, feedCard
	msgtype string       `json:"msgtype"`
	text    NullableText `json:"text,omitempty"`
}

// NewMessageWith instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageWith(msgtype string) *Message {
	this := Message{}
	this.msgtype = msgtype
	return &this
}

// NewMessage instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessage() *Message {
	this := Message{}
	var msgtype string = "text"
	this.msgtype = msgtype
	return &this
}

// GetActionCard returns the ActionCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetActionCard() ActionCard {
	if o == nil || utils.IsNil(o.actionCard.Get()) {
		var ret ActionCard
		return ret
	}
	return *o.actionCard.Get()
}

// GetActionCardOk returns a tuple with the ActionCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetActionCardOk() (*ActionCard, bool) {
	if o == nil {
		return nil, false
	}
	return o.actionCard.Get(), o.actionCard.IsSet()
}

// HasActionCard returns a boolean if a field has been set.
func (o *Message) HasActionCard() bool {
	if o != nil && o.actionCard.IsSet() {
		return true
	}

	return false
}

// ActionCard gets a reference to the given NullableActionCard and assigns it to the actionCard field.
// actionCard:
func (o *Message) ActionCard(v ActionCard) *Message {
	o.actionCard.Set(&v)
	return o
}

// SetActionCardNil sets the value for ActionCard to be an explicit nil
func (o *Message) SetActionCardNil() {
	o.actionCard.Set(nil)
}

// UnsetActionCard ensures that no value is present for ActionCard, not even an explicit nil
func (o *Message) UnsetActionCard() {
	o.actionCard.Unset()
}

// GetAt returns the At field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetAt() At {
	if o == nil || utils.IsNil(o.at.Get()) {
		var ret At
		return ret
	}
	return *o.at.Get()
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetAtOk() (*At, bool) {
	if o == nil {
		return nil, false
	}
	return o.at.Get(), o.at.IsSet()
}

// HasAt returns a boolean if a field has been set.
func (o *Message) HasAt() bool {
	if o != nil && o.at.IsSet() {
		return true
	}

	return false
}

// At gets a reference to the given NullableAt and assigns it to the at field.
// at:
func (o *Message) At(v At) *Message {
	o.at.Set(&v)
	return o
}

// SetAtNil sets the value for At to be an explicit nil
func (o *Message) SetAtNil() {
	o.at.Set(nil)
}

// UnsetAt ensures that no value is present for At, not even an explicit nil
func (o *Message) UnsetAt() {
	o.at.Unset()
}

// GetFeedCard returns the FeedCard field value if set, zero value otherwise.
func (o *Message) GetFeedCard() FeedCard {
	if o == nil || utils.IsNil(o.feedCard) {
		var ret FeedCard
		return ret
	}
	return *o.feedCard
}

// GetFeedCardOk returns a tuple with the FeedCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetFeedCardOk() (*FeedCard, bool) {
	if o == nil || utils.IsNil(o.feedCard) {
		return nil, false
	}
	return o.feedCard, true
}

// HasFeedCard returns a boolean if a field has been set.
func (o *Message) HasFeedCard() bool {
	if o != nil && !utils.IsNil(o.feedCard) {
		return true
	}

	return false
}

// FeedCard gets a reference to the given FeedCard and assigns it to the feedCard field.
// feedCard:
func (o *Message) FeedCard(v FeedCard) *Message {
	o.feedCard = &v
	return o
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Message) GetLink() Link {
	if o == nil || utils.IsNil(o.link) {
		var ret Link
		return ret
	}
	return *o.link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetLinkOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.link) {
		return nil, false
	}
	return o.link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Message) HasLink() bool {
	if o != nil && !utils.IsNil(o.link) {
		return true
	}

	return false
}

// Link gets a reference to the given Link and assigns it to the link field.
// link:
func (o *Message) Link(v Link) *Message {
	o.link = &v
	return o
}

// GetMarkdown returns the Markdown field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetMarkdown() Markdown {
	if o == nil || utils.IsNil(o.markdown.Get()) {
		var ret Markdown
		return ret
	}
	return *o.markdown.Get()
}

// GetMarkdownOk returns a tuple with the Markdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetMarkdownOk() (*Markdown, bool) {
	if o == nil {
		return nil, false
	}
	return o.markdown.Get(), o.markdown.IsSet()
}

// HasMarkdown returns a boolean if a field has been set.
func (o *Message) HasMarkdown() bool {
	if o != nil && o.markdown.IsSet() {
		return true
	}

	return false
}

// Markdown gets a reference to the given NullableMarkdown and assigns it to the markdown field.
// markdown:
func (o *Message) Markdown(v Markdown) *Message {
	o.markdown.Set(&v)
	return o
}

// SetMarkdownNil sets the value for Markdown to be an explicit nil
func (o *Message) SetMarkdownNil() {
	o.markdown.Set(nil)
}

// UnsetMarkdown ensures that no value is present for Markdown, not even an explicit nil
func (o *Message) UnsetMarkdown() {
	o.markdown.Unset()
}

// GetMsgtype returns the Msgtype field value
func (o *Message) GetMsgtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.msgtype
}

// GetMsgtypeOk returns a tuple with the Msgtype field value
// and a boolean to check if the value has been set.
func (o *Message) GetMsgtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.msgtype, true
}

// Msgtype sets field value
func (o *Message) Msgtype(v string) *Message {
	o.msgtype = v
	return o
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Message) GetText() Text {
	if o == nil || utils.IsNil(o.text.Get()) {
		var ret Text
		return ret
	}
	return *o.text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Message) GetTextOk() (*Text, bool) {
	if o == nil {
		return nil, false
	}
	return o.text.Get(), o.text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *Message) HasText() bool {
	if o != nil && o.text.IsSet() {
		return true
	}

	return false
}

// Text gets a reference to the given NullableText and assigns it to the text field.
// text:
func (o *Message) Text(v Text) *Message {
	o.text.Set(&v)
	return o
}

// SetTextNil sets the value for Text to be an explicit nil
func (o *Message) SetTextNil() {
	o.text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *Message) UnsetText() {
	o.text.Unset()
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.actionCard.IsSet() {
		toSerialize["actionCard"] = o.actionCard.Get()
	}
	if o.at.IsSet() {
		toSerialize["at"] = o.at.Get()
	}
	if !utils.IsNil(o.feedCard) {
		toSerialize["feedCard"] = o.feedCard
	}
	if !utils.IsNil(o.link) {
		toSerialize["link"] = o.link
	}
	if o.markdown.IsSet() {
		toSerialize["markdown"] = o.markdown.Get()
	}
	toSerialize["msgtype"] = o.msgtype
	if o.text.IsSet() {
		toSerialize["text"] = o.text.Get()
	}
	return toSerialize, nil
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
